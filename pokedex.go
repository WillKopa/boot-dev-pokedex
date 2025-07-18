package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/WillKopa/boot_dev_pokedex/pokecache"
)

func cleanInput(text string) []string {
	strings_slice := []string{}
	text = strings.TrimSpace(text)
	for _, str := range(strings.Fields(text)) {
		str = strings.TrimSpace(str)
		str = strings.ToLower(str)
		strings_slice = append(strings_slice, str)
	}

	return strings_slice
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	base_url := "https://pokeapi.co/api/v2/location-area/"
	cache := pokecache.NewCache(5 * time.Second)
	c := &config{
		Next: &base_url,
		Previous: nil,
		Cache: cache,
	}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		user_input := scanner.Text()
		input_slice := cleanInput(user_input)
		if len(input_slice) != 0 {
			handleCommand(input_slice[0], c)
		}
	}
}