DB_CONTAINER_NAME = postgresql

run:
	APP_ENV=local go run cmd/weather-api/main.go
build:
	go build -o weather-api cmd/weather-api/main.go
docker-build:
	docker build --build-arg APP_ENV=dev -t weather-api .
run-tests:
	go test -count=1 ./...
lint:
	gofmt -w .
	goimports -w .
	golangci-lint run -v
generate-mocks:
	mockgen -source=internal/weather/service.go -destination=internal/weather/mock/service_mock.go -package=mocks
	mockgen -source=internal/weather/repository.go -destination=internal/weather/mock/repository_mock.go -package=mocks
ready-entire-system:
	docker-compose up -d
	sleep 5
	make shut-logging
	make install-http-extension
	make cp-files
	make run-files
install-http-extension:
	docker cp .db/install-http-extension.sh $(DB_CONTAINER_NAME):/tmp/install-http-extension.sh
	docker exec -i $(DB_CONTAINER_NAME) bash -c "chmod +x /tmp/install-http-extension.sh && /tmp/install-http-extension.sh"
shut-logging:
	docker cp .db/shut-logging.sh $(DB_CONTAINER_NAME):/tmp/shut-logging.sh
	docker exec -i $(DB_CONTAINER_NAME) bash -c "chmod +x /tmp/shut-logging.sh && /tmp/shut-logging.sh"
	docker restart $(DB_CONTAINER_NAME)
cp-files:
	docker cp .db/queries/create_table.sql $(DB_CONTAINER_NAME):/create_table.sql
	docker cp .db/functions/notify_users.sql $(DB_CONTAINER_NAME):/notify_users.sql
	docker cp .db/functions/update_weather.sql $(DB_CONTAINER_NAME):/update_weather.sql
	docker cp .db/functions/get_temperatures.sql $(DB_CONTAINER_NAME):/get_temperatures.sql
	docker cp .db/functions/get_present_first_request_uuid.sql $(DB_CONTAINER_NAME):/get_present_first_request_uuid.sql
	docker cp .db/triggers/check_request_count.sql $(DB_CONTAINER_NAME):/check_request_count.sql
	docker cp .db/triggers/check_first_request.sql $(DB_CONTAINER_NAME):/check_first_request.sql
	docker cp .db/create_triggers.sql $(DB_CONTAINER_NAME):/create_triggers.sql
	docker cp .db/create_index_for_location.sql $(DB_CONTAINER_NAME):/create_index_for_location.sql
run-files:
	docker exec -it $(DB_CONTAINER_NAME) psql -U admin -d weather -a \
		-f create_table.sql \
		-f notify_users.sql \
		-f update_weather.sql \
		-f get_temperatures.sql \
		-f get_present_first_request_uuid.sql \
		-f check_request_count.sql \
		-f check_first_request.sql \
		-f create_triggers.sql \
		-f create_index_for_location.sql
