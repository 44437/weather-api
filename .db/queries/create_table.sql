CREATE TABLE "weather_queries" (
  "id" SERIAL PRIMARY KEY,
  "location" TEXT NOT NULL,
  "service_1_temperature" TEXT NULL,
  "service_2_temperature" TEXT NULL,
  "request_count" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "first_request_uuid" VARCHAR(36) NOT NULL
  );