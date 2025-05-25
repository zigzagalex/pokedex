package commands

import (
	"fmt"
	"os"
)

func CommandExit(conf *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
