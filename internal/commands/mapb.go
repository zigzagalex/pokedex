package commands

import (
	"fmt"

	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

func commandMapBack(conf *Config) error {
	url := conf.Prev
	if url == "" {
		fmt.Println("You're on the first page.")
		return nil
	}
	data, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	conf.Next = data.Next
	conf.Prev = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
