package main

import "fmt"

func help() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, i := range commands {
		fmt.Printf("%s: %s\n", i.name, i.description)
	}
	return nil
}
