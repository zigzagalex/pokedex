package commands

import (
	"encoding/json"
	"fmt"

	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

func CommandMapBack(conf *Config, args ...string) error {
	url := conf.Prev
	if url == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	var body []byte
	var err error

	body, ok := conf.Cache.Get(url)
	if !ok {
		body, err = pokeapi.GetResult(url)
		if err != nil {
			return err
		}
		conf.Cache.Add(url, body)
	}

	var data pokeapi.PokeAPIResult
	err = json.Unmarshal(body, &data)
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
