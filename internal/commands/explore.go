package commands

import (
	"encoding/json"
	"fmt"

	"github.com/zigzagalex/pokedex/internal/pokeapi"
)

type LocationAreaDetails struct {
	Name                 string                `json:"name"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type EncounterDetail struct {
	Chance   int `json:"chance"`
	MinLevel int `json:"min_level"`
	MaxLevel int `json:"max_level"`
	Method   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"method"`
	ConditionValues []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"condition_values"`
}

type EncounterMethodRate struct {
	EncounterMethod struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"encounter_method"`
	VersionDetails []struct {
		Rate    int `json:"rate"`
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}

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
