package main

import "fmt"

func commandPokedex(conf *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	if len(conf.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet!")
		return nil
	}
	for name, _ := range conf.caughtPokemon {
		fmt.Println("\t- ", name)
	}

	return nil
}
