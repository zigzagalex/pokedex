package commands

import (
	"encoding/json"
	"fmt"

	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

func CommandExplore(conf *Config, args ...string) error {
	if args[0] == "" {
		fmt.Println("Explore command needs a loacation area, e.g: explore oreburgh-gate-1f")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

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

	var data LocationAreaDetails
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", data)

	fmt.Printf("Exploring %s...\n", data.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
