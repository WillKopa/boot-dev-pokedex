package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	pokemon_api "github.com/WillKopa/boot_dev_pokedex/api"
	"github.com/WillKopa/boot_dev_pokedex/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Pokedex           map[string]pokemon_api.Pokemon
	Args              []string
	Base_pokemon_url  *string
	Base_location_url *string
	Next              *string
	Previous          *string
	Cache             *pokecache.Cache
}

func handleCommand(commandName string, c *config) {
	command, exists := getCommands()[commandName]
	if exists {
		err := command.callback(c)
		if err != nil {
			fmt.Printf("Error handling command, %v\n", err)
		}
	} else {
		fmt.Println("Unkown command")
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help menu",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations in the Pokemon world. Calling map again will display the next 20 in the list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world. Repeated calls with continue to call the previous results",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Takes a location name and returns a list of pokemon that are in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon with catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you have caught using inspect <pokemon>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemon that have been caught so far",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(_ *config) error {
	help_message := "Welcome to the Pokedex!\nUsage:\n"

	for _, command := range getCommands() {
		help_message += fmt.Sprintf("\n%s: %s", command.name, command.description)
	}

	fmt.Println(help_message)
	return nil
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Program didn't exit")
}

func commandMap(c *config) error {
	if c.Next == nil {
		return fmt.Errorf("end of map reached. cannot progress further")
	}

	locations, err := pokemon_api.GetLocationsFromAPI(c.Next, c.Cache)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous
	err = printMap(locations)
	if err != nil {
		return err
	}
	return nil
}

func commandMapBack(c *config) error {
	if c.Previous == nil {
		return fmt.Errorf("beginning of map reached. cannot go back farther")
	}

	locations, err := pokemon_api.GetLocationsFromAPI(c.Previous, c.Cache)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	err = printMap(locations)
	if err != nil {
		return err
	}
	return nil
}

func printMap(l pokemon_api.Locations_api_response) error {
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(c *config) error {
	area := strings.Join(c.Args, "")
	fmt.Println("Exploring " + area + "...")
	full_url := *c.Base_location_url + area
	pokemon, err := pokemon_api.GetPokemonInLocationFromAPI(&full_url, c.Cache)
	if err != nil {
		return fmt.Errorf("error calling api: %v", err)
	}

	err = printPokemonInArea(pokemon)
	if err != nil {
		return fmt.Errorf("error printing list of pokemon: %v", err)
	}
	return nil
}

func printPokemonInArea(p pokemon_api.Pokemon_in_location_response) error {
	encounters := p.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config) error {
	pokemon_name := strings.Join(c.Args, "")
	full_url := *c.Base_pokemon_url + pokemon_name
	fmt.Println("Throwing a Pokeball at " + pokemon_name + "...")
	pokemon, err := pokemon_api.GetPokemonFromAPI(&full_url, c.Cache)
	if err != nil {
		return fmt.Errorf("error getting pokemon from api: %v", err)
	}
	if rand.Intn(pokemon.Base_experience) < 50 {
		c.Pokedex[pokemon.Name] = pokemon
		fmt.Println(pokemon.Name + " was caught!")
		fmt.Println("You can now inspect the pokemon with the inspect command")
	} else {
		fmt.Println(pokemon_name + " escaped!")
	}
	return nil
}

func commandInspect(c *config) error {
	pokemon_name := strings.Join(c.Args, "")
	poke_stats, exists := c.Pokedex[pokemon_name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", poke_stats.Name)
	fmt.Printf("Height: %v\n", poke_stats.Height)
	fmt.Printf("Weight: %v\n", poke_stats.Weight)
	fmt.Println("Stats: ")
	for _, stat := range(poke_stats.Stats) {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poke_type := range(poke_stats.Types) {
		fmt.Println("  - " + poke_type.Type.Name)
	}

	return nil
}

func commandPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	if len(c.Pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
		fmt.Println("Catch pokemon by using the catch command")
	} else {
		for _, pokemon := range(c.Pokedex) {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}
