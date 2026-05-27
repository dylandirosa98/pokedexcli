package main

import "fmt"

func pokedex(fig *config, arg ...string) error {
	print("Your Pokedex:\n")
	for pokemon, _ := range fig.Pokedex {
		fmt.Printf("  - %s\n", pokemon)
	}
	return nil
}
