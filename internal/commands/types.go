package commands

type Config struct {
	Prev string
	Next string
}

type CLICommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}