package main

import (
	"fmt"
	"os"
)

func exit(fig *config, arg ...string) error {
	print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("exiting the program")
}
