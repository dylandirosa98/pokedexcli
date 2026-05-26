package main

import (
	"fmt"
	"os"
)

func exit(fig *config) error {
	print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("exiting the program")
}
