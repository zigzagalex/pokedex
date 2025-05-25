package commands

import "github.com/zigzagalex/pokedex/internal/pokecache"

type Config struct {
	Prev    string
	Next    string
	Cache   *pokecache.Cache
	Pokedex map[string]Pokemon
}

type CLICommand struct {
	Name        string
	Description string
	Callback    func(conf *Config, args ...string) error
}

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

type Pokemon struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	IsDefault      bool        `json:"is_default"`
	Order          int         `json:"order"`
	Weight         int         `json:"weight"`
	Abilities      []Ability   `json:"abilities"`
	Types          []TypeSlot  `json:"types"`
	Stats          []StatEntry `json:"stats"`
}

type Ability struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
}

type TypeSlot struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type StatEntry struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}
