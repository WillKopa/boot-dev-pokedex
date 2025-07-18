package pokemon_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WillKopa/boot_dev_pokedex/pokecache"
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

type Pokemon_in_location_response struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationsFromAPI(url *string, cache *pokecache.Cache) (Locations_api_response, error) {
	body, exists := cache.Get(*url)
	if !exists {
		res, err := http.Get(*url)
		if err != nil {
			return Locations_api_response{}, fmt.Errorf("error calling area endpoint: %v", err)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Locations_api_response{}, fmt.Errorf("response code %v when calling area endpoint", res.StatusCode)
		}

		if err != nil {
			return Locations_api_response{}, fmt.Errorf("error when reading response body: %v", err)
		}
	}
	cache.Add(*url, body)
	locations := Locations_api_response{}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return Locations_api_response{}, fmt.Errorf("error Unmarshaling body to Locations_api_response struct: %v", err)
	}
	return locations, nil
}

func GetPokemonInLocationFromAPI(url *string, cache *pokecache.Cache) (Pokemon_in_location_response, error) {
	body, exists := cache.Get(*url)
	if !exists {
		res, err := http.Get(*url)
		if err != nil {
			return Pokemon_in_location_response{}, fmt.Errorf("error calling area endpoint: %v", err)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Pokemon_in_location_response{}, fmt.Errorf("response code %v when calling area endpoint", res.StatusCode)
		}

		if err != nil {
			return Pokemon_in_location_response{}, fmt.Errorf("error when reading response body: %v", err)
		}
	}
	cache.Add(*url, body)
	pokemon := Pokemon_in_location_response{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon_in_location_response{}, fmt.Errorf("error Unmarshaling body to Pokemon_in_location_response struct: %v", err)
	}
	return pokemon, nil
}