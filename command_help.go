package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range conf.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
