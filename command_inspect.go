package main

import "fmt"

func inspect(fig *config, arg ...string) error {
	pokemon, ok := fig.Pokedex[arg[0]]
	if ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		stats := pokemon.Stats
		for _, stat := range stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		print("Types:\n")
		types := pokemon.Types
		for _, theType := range types {
			fmt.Printf("  - %s\n", theType.Type.Name)
		}
		return nil
	}
	print("you have not caught that pokemon\n")
	return fmt.Errorf("havent caught the pokemon")
}
