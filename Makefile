run:
	APP_ENV=local go run cmd/weather-api/main.go
build:
	go build -o weather-api cmd/weather-api/main.go
docker-build:
	docker build --build-arg APP_ENV=dev -t weather-api .
lint:
	gofmt -w .
	goimports -w .
	golangci-lint run -v
generate-mocks:
	mockgen -source=internal/weather/service.go -destination=internal/weather/mock/service_mock.go -package=mocks
	mockgen -source=internal/weather/repository.go -destination=internal/weather/mock/repository_mock.go -package=mocks
