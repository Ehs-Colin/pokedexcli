package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	// 1> Ensure a name has been entered
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}
	pokemonName := args[0]
	// 2> Check if pokemon is listed in caught pokemon
	pokemon, exists := conf.caughtPokemon[pokemonName]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("\t- %s\n", pokemonType.Type.Name)
	}
	return nil
}
