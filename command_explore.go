package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	locationName := args[0]
	location, err := conf.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
