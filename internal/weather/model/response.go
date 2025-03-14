package model

type Response struct {
	Location    string  `json:"location"`
	Temperature float32 `json:"temperature"`
}
