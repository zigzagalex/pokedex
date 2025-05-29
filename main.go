package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/peterh/liner"
	"github.com/zigzagalex/pokedex/internal/commands"
	"github.com/zigzagalex/pokedex/internal/pokecache"
)

var (
	pokedex_history = filepath.Join(os.TempDir(), ".pokedex_history")
	names           = []string{"explore", "inspect", "pokedex"}
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

	// Start Liner
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	// Set cache file
	historyFile := ".pokedex_history"
	if f, err := os.Open(historyFile); err == nil {
		defer f.Close()
		line.ReadHistory(f)
	}

	for {
		input_text, err := line.Prompt("Pokedex > ")
		if err != nil {
			break
		}
		if input_text == "" {
			continue
		}
		input_text = strings.TrimSpace(input_text)
		line.AppendHistory(input_text)

		tokens := cleanInput(input_text)
		cmdName := tokens[0]

		var args string
		if len(tokens) > 1 {
			args = tokens[1]
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

	}
	f, err := os.Create(historyFile)
	if err == nil {
		defer f.Close()
		line.WriteHistory(f)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
