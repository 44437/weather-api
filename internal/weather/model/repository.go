package model

import "time"

type DBWeather struct {
	ID                  int       `json:"id"`
	Location            string    `json:"location"`
	Service1Temperature *float32  `json:"service_1_temperature"`
	Service2Temperature *float32  `json:"service_2_temperature"`
	RequestCount        uint8     `json:"request_count"`
	CreatedAt           time.Time `json:"created_at"`
	FirstRequestUUID    string    `json:"first_request_uuid"`
}
