package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

func CommandCatch(conf *Config, args ...string) error {
	name := args[0]
	if name == "" {
		fmt.Println("To catch a Pokemon you need to use: catch <name>")
		return nil
	}
	if _, alreadyCaught := conf.Pokedex[name]; alreadyCaught {
		fmt.Printf("You already caught %s! It's in your Pokédex.\n", name)
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + name

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

	var poke Pokemon
	err = json.Unmarshal(body, &poke)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", poke.Name)

	base := float64(poke.BaseExperience)
	weight := float64(poke.Weight)

	difficulty := (base * 0.8) + (weight * 0.2)
	maxDifficulty := 200.0
	catchChance := 1.0 - (difficulty / maxDifficulty)
	if catchChance < 0.1 {
		catchChance = 0.1
	}

	roll := rand.Float64()

	if roll < catchChance {
		fmt.Printf("%s was caught!\n", poke.Name)
		conf.Pokedex[poke.Name] = poke
	} else {
		fmt.Printf("%s escaped!\n", poke.Name)
	}

	return nil
}
