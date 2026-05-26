package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
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
			commands[text[0]].callback()
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
