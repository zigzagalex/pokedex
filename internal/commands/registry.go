package commands

func AvailableCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"easteregg": {
			Name:        "easteregg",
			Description: "Surprise easter-egg",
			Callback:    CommandEasteregg,
		},
		"map": {
			Name:        "map",
			Description: "Shows next 20 map areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows previous 20 map areas in the Pokemon world",
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Lets you explore a location area with explore <location-area>",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "User can try and catch a Pokemon with catch <name>",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Shows the caught pokemon",
			Callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Get info of caught pokemon",
			Callback:    CommandInspect,
		},
	}
}
