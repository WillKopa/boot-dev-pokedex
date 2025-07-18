package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/WillKopa/boot_dev_pokedex/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
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
	locations, err := pokemon_api.GetLocationsFromAPI(c.Next)
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

	locations, err := pokemon_api.GetLocationsFromAPI(c.Previous)
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
