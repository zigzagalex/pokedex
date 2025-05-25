package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/zigzagalex/pokedex/internal/commands"
	"github.com/zigzagalex/pokedex/internal/pokecache"
)

func main() {
	cmds := commands.AvailableCommands()
	cache, err := pokecache.NewCache(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	conf := commands.Config{
		Cache:   &cache,
		Pokedex: make(map[string]commands.Pokemon),
	}

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		cmdName := text[0]
		var args string
		if len(text) > 1 {
			args = text[1]
		}
		command, ok := cmds[cmdName]
		if ok {
			err := command.Callback(&conf, args)
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
