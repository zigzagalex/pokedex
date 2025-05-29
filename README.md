# pokedex
A cli REPL using the PokeAPI.

The goal of the project is to learn/practice the following:
1. Creating a working go module
2. Interfacing with an api
3. Creating cache
4. Building a simple REPL


When working with an API it is helpful to create an API directory that will contain any generic interaction with the API. This helps to separate out concerns. 

A CLI as a simple REPL must manage inputs, transform them and give an output. A simple way is to use the liner module. (https://github.com/peterh/liner)

Here is a basic skeleton for a CLI REPL: 

```
package main

import(
  "fmt"
  "github.com/peterh/liner"
)

var (
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
	names      = []string{"string1", "something2", "anything3", "nothing4"}
)

func main {
  // Preoperations

  // Start the REPL using liner
  line := liner.NewLiner()
  defer line.Close()
  // Define liner settings
  line.SetCtrlCAborts(true)
  // Set completer function 
  line.SetCompleter(func(line string) (c []string) {
    for _, n := range names {
	  if strings.HasPrefix(n,strings.ToLower(line)) {
		c = append(c, n)
		}
	}
	return
  })
  // Set cache file
  if f, err := os.Open(history_fn); err == nil {
	  defer f.Close()
	  line.ReadHistory(f)
  }

  // Set REPL loop
  for {
    input_text, err := line.Prompt("<PROMPT_TEXT>")
	if err != nil {
	  break 
	}
	if input_text == "" {
	  continue
	}
	input_text = strings.TrimSpace(input_text)
	line.AppendHistory(input_text)

    // Evaluate the input_text (functionality lives here)

  } // loop end
  // save cached commands
  f, err := os.Create(historyFile)
	if err == nil {
		defer f.Close()
		line.WriteHistory(f)
}
```

A CLI REPL is a very easy way to build an MVP, since you skip all the UX design. I think it is a good strategy to setup a CLI and make sure the backbone of functionality is set up correctly and then deal with the UX, because that will then be a mapping of commands/functionality to graphical elements. 

Ways to extend the project: 
* Simulate battles between pokemon
* Allow multiple players to play at the same time
* Allow for pokemon that are caught to evolve after a set amount of time
* Persist a user's Pokedex to disk so they can save progress between sessions
* Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
* Random encounters with wild pokemon
* Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
