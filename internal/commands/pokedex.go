package commands

import "fmt"

func CommandPokedex(conf *Config, args ...string) error {
	if len(conf.Pokedex) == 0 {
		fmt.Println("You haven’t caught any Pokémon yet.")
		return nil
	}

	fmt.Println("Your Pokédex:")
	for name := range conf.Pokedex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
