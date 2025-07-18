package pokemon_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Locations_api_response struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Location
}

func GetLocationsFromAPI(url *string) (Locations_api_response, error) {
	res, err := http.Get(*url)
	if err != nil {
		return Locations_api_response{}, fmt.Errorf("error calling area endpoint: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Locations_api_response{}, fmt.Errorf("response code %v when calling area endpoint", res.StatusCode)
	}

	if err != nil {
		return Locations_api_response{}, fmt.Errorf("error when reading response body: %v", err)
	}

	locations := Locations_api_response{}

	err = json.Unmarshal(body, &locations)

	if err != nil {
		return Locations_api_response{}, fmt.Errorf("error Unmarshaling body to Locations_api_response struct: %v", err)
	}
	return locations, nil
}
