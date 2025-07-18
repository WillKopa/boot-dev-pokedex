package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		user_input := scanner.Text()
		input_slice := cleanInput(user_input)
		if len(input_slice) != 0 {
			fmt.Printf("Your command was: %v\n", input_slice[0])
			if input_slice[0] == "exit" {
				os.Exit(0)
			}
		}
	}
}