package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandEasteregg() error {
	fmt.Println("🙉")
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"easteregg": {
			name:        "easteregg",
			description: "Surprise easter-egg",
			callback:    commandEasteregg,
		},
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}
		command, ok := commands[text[0]]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Command not found.")
		}
		fmt.Print("Pokedex > ")
	}

}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}
