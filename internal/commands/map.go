package commands

import (
	"fmt"
	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

func CommandMap(conf *Config) error {
	url := conf.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
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
