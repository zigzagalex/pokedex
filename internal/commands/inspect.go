package commands

import (
	"fmt"
)

func CommandInspect(conf *Config, args ...string) error {
	name := args[0]
	if poke, alreadyCaught := conf.Pokedex[name]; alreadyCaught {
		fmt.Println(poke)
		return nil
	} else {
		fmt.Printf("You haven't caught %v yet.\n", name)
	}
	return nil
}

func (p Pokemon) String() string {
	s := fmt.Sprintf(
		"\nName:           %s\nID:             %d\nBase XP:        %d\nHeight:         %d\nWeight:         %d\nOrder:          %d\nIs Default Form: %v\n",
		p.Name,
		p.ID,
		p.BaseExperience,
		p.Height,
		p.Weight,
		p.Order,
		p.IsDefault,
	)

	// Print Types
	if len(p.Types) > 0 {
		s += "Types:          "
		for i, t := range p.Types {
			if i > 0 {
				s += ", "
			}
			s += t.Type.Name
		}
		s += "\n"
	}

	// Print Abilities
	if len(p.Abilities) > 0 {
		s += "Abilities:      "
		for i, a := range p.Abilities {
			if i > 0 {
				s += ", "
			}
			if a.IsHidden {
				s += fmt.Sprintf("%s (Hidden)", a.Ability.Name)
			} else {
				s += a.Ability.Name
			}
		}
		s += "\n"
	}

	// Print Stats
	if len(p.Stats) > 0 {
		s += "Stats:\n"
		for _, stat := range p.Stats {
			s += fmt.Sprintf("  - %-12s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}

	return s
}
