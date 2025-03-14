FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o weather-api cmd/weather-api/main.go

FROM alpine:3.21.3 AS prod
WORKDIR /app/

ARG APP_ENV
ENV APP_ENV=$APP_ENV

COPY --from=builder /app/weather-api ./
COPY ./.config ./.config

RUN chmod +x weather-api

EXPOSE 8080
CMD [ "./weather-api" ]