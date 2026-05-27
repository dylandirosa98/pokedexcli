package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dylandirosa98/pokedexcli/internal"
)

func cleanInput(text string) []string {
	newText := strings.ToLower(text)
	words := strings.Fields(newText)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
		"help": {
			name:        "help",
			description: "Displays help mesage",
			callback:    help,
		},
		"map": {
			name:        "map",
			description: "displays location areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "displays previous location areas",
			callback:    mapb,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	time := 5 * time.Second
	cache := internal.NewCache(time)
	con := config{
		nextURL:     nil,
		previousURL: nil,
		cache:       &cache,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		if commands[text[0]].name == "" {
			print("Unknown command")
		} else {
			commands[text[0]].callback(&con)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL     *string
	previousURL *string
	cache       *internal.Cache
}
