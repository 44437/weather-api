FROM golang:1.24-alpine

ARG APP_ENV
ENV APP_ENV=$APP_ENV

WORKDIR /app/weather-api
COPY ./weather-api ./

RUN apk add --no-cache libc6-compat
RUN chmod +x weather-api

EXPOSE 8080
CMD [ "./weather-api" ]