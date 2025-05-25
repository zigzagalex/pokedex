package commands

import "github.com/zigzagalex/pokedex/internal/pokecache"


type Config struct {
	Prev string
	Next string
	Cache *pokecache.Cache
}

type CLICommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}