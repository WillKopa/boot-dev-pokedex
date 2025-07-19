package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/WillKopa/boot_dev_pokedex/pokecache"
	"github.com/WillKopa/boot_dev_pokedex/api"
)

func cleanInput(text string) []string {
	strings_slice := []string{}
	text = strings.TrimSpace(text)
	for _, str := range strings.Fields(text) {
		str = strings.TrimSpace(str)
		str = strings.ToLower(str)
		strings_slice = append(strings_slice, str)
	}

	return strings_slice
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	base_location_url := "https://pokeapi.co/api/v2/location-area/"
	base_pokemon_url := "https://pokeapi.co/api/v2/pokemon/"
	cache := pokecache.NewCache(5 * time.Second)
	c := &config{
		Pokedex:		   map[string]pokemon_api.Pokemon{},
		Base_pokemon_url:  &base_pokemon_url,
		Base_location_url: &base_location_url,
		Next:              &base_location_url,
		Previous:          nil,
		Cache:             cache,
	}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		user_input := scanner.Text()
		input_slice := cleanInput(user_input)
		if len(input_slice) != 0 {
			c.Args = input_slice[1:]
			handleCommand(input_slice[0], c)
		}
	}
}
