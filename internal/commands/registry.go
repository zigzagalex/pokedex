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
			Callback:    commandMapBack,
		},
	}
}
