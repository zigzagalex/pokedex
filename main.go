package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zigzagalex/pokedex/internal/commands"
)

func main() {
	conf := commands.Config{}
	cmds := commands.AvailableCommands()

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		command, ok := cmds[text[0]]
		if ok {
			err := command.Callback(&conf)
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
	return strings.Fields(strings.ToLower(text))
}
