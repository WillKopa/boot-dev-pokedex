package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locations_api_response struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []location
}

func getLocationsFromAPI(url *string) (locations_api_response, error) {
	res, err := http.Get(*url)
	if err != nil {
		return locations_api_response{}, fmt.Errorf("error calling area endpoint: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return locations_api_response{}, fmt.Errorf("response code %v when calling area endpoint", res.StatusCode)
	}

	if err != nil {
		return locations_api_response{}, fmt.Errorf("error when reading response body: %v", err)
	}

	locations := locations_api_response{}

	err = json.Unmarshal(body, &locations)

	if err != nil {
		return locations_api_response{}, fmt.Errorf("error Unmarshaling body to locations_api_response struct: %v", err)
	}
	return locations, nil
}
