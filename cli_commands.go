package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name 			string
	description 	string
	callback		func() error
}

func handleCommand(commandName string) {
	command, exists := getCommands()[commandName]
	if exists {
		command.callback()
	} else {
		fmt.Println("Unkown command")
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:			"exit",
			description:	"Exit the pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:			"help",
			description: 	"Display help menu",
			callback:		commandHelp,
		},
	}
}

func commandHelp() error {
	help_message := "Welcome to the Pokedex!\nUsage:\n"

	for _, command := range(getCommands()) {
		help_message += fmt.Sprintf("\n%s: %s", command.name, command.description)
	}

	fmt.Println(help_message)
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Program didn't exit")
}