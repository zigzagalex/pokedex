package commands

import "github.com/zigzagalex/pokedex/internal/pokecache"


type Config struct {
	Prev string
	Next string
	Cache *pokecache.Cache
	Pokedex map[string]Pokemon
}

type CLICommand struct {
	Name        string
	Description string
	Callback    func(conf *Config, args ...string) error
}