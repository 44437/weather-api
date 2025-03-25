CREATE TABLE "weather_queries" (
  "id" SERIAL PRIMARY KEY,
  "location" TEXT NOT NULL,
  "service_1_temperature" FLOAT NULL,
  "service_2_temperature" FLOAT NULL,
  "request_count" SMALLINT NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "first_request_uuid" UUID NOT NULL
  );