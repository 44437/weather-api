run:
	APP_ENV=local go run cmd/weather-api/main.go

lint:
	gofmt -w .
	goimports -w .
	golangci-lint run -v