package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil //because of compiler complaints.  Seriously, it exits the program a line above!
}
