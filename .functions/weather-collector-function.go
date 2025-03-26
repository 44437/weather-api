package functions

// It has been manually deployed on gcp for the moment
// curl -X GET https://weather-collector-function.ercant.net?location=<location>
// response: {"service_1_temperature":18.1,"service_2_temperature":18}

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetTemperatures", getTemperatures)
}

type Service1Response struct {
	Current Current1                `json:"current"`
	Error   *map[string]interface{} `json:"error"`
}

type Current1 struct {
	Temperature float32 `json:"temp_c"`
}

type Service2Response struct {
	Current Current2                `json:"current"`
	Error   *map[string]interface{} `json:"error"`
}

type Current2 struct {
	Temperature float32 `json:"temperature"`
}

type Response interface{}

type GeneralResponse struct {
	Response Response
	Error    interface{}
}

type OutgoingResponse struct {
	Service1Temperature float32 `json:"service_1_temperature"`
	Service2Temperature float32 `json:"service_2_temperature"`
}

func getWeatherFromService1(location string, ch chan<- GeneralResponse) {
	defer func() {
		if r := recover(); r != nil {
			ch <- GeneralResponse{
				Response: nil,
				Error:    r,
			}
		}
	}()

	url := fmt.Sprintf("weatherapi-url", location)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Panicf("Error whilst making a GET request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	var response Service1Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil || response.Error != nil {
		log.Panicf("Error whilst decoding JSON response from %s: %v", url, err)
	}

	ch <- GeneralResponse{
		Response: response,
		Error:    nil,
	}
}

func getWeatherFromService2(location string, ch chan<- GeneralResponse) {
	defer func() {
		if r := recover(); r != nil {
			ch <- GeneralResponse{
				Response: nil,
				Error:    r,
			}
		}
	}()

	url := fmt.Sprintf("weatherstack-url", location)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Panicf("Error whilst making a GET request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	var response Service2Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil || response.Error != nil {
		log.Panicf("Error whilst decoding JSON response from %s: %v", url, err)
	}

	ch <- GeneralResponse{
		Response: response,
		Error:    nil,
	}
}

func getTemperatures(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ch1 := make(chan GeneralResponse)
	ch2 := make(chan GeneralResponse)

	go getWeatherFromService1(location, ch1)

	go getWeatherFromService2(location, ch2)

	response1 := <-ch1
	response2 := <-ch2

	if response1.Error != nil || response2.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := OutgoingResponse{
		Service1Temperature: response1.Response.(Service1Response).Current.Temperature,
		Service2Temperature: response2.Response.(Service2Response).Current.Temperature,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Unable to encode data: %v", err), http.StatusInternalServerError)
		return
	}
}
