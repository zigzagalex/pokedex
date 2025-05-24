package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	Prev string
	Next string
}

type pokeAPIResult struct {
	Count    int
	Next     string
	Previous string
	Results  []locationArea
}

type locationArea struct {
	Name string
	URL  string
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandEasteregg(conf *config) error {
	fmt.Println("🙉")
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *config) error {
	var lookupURL string
	if conf.Next == "" {
		lookupURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		lookupURL = conf.Next
	}
	res, err0 := http.Get(lookupURL)
	if err0 != nil {
		fmt.Println(err0)
	}
	defer res.Body.Close()

	var data pokeAPIResult
	err1 := json.NewDecoder(res.Body).Decode(&data)
	if err1 != nil {
		fmt.Println(err1)
	}
	conf.Prev = data.Previous
	conf.Next = data.Next
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

var commands map[string]cliCommand
var conf config

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
		"map": {
			name:        "map",
			description: "Shows 20 map areas in the Pokemon world",
			callback:    commandMap,
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
			err := command.callback(&conf)
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
