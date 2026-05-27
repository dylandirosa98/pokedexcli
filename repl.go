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
		"explore": {
			name:        "explore",
			description: "explore an area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "see details about a pokemon",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "view all of the pokemon you caught",
			callback:    pokedex,
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
		Pokedex:     make(map[string]Pokemon),
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
			if len(text) > 1 {
				commands[text[0]].callback(&con, text[1])
			} else {
				commands[text[0]].callback(&con)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextURL     *string
	previousURL *string
	cache       *internal.Cache
	Pokedex     map[string]Pokemon
}
