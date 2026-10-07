package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	// 1> Ensure a name has been entered
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}
	// 2> Use api call to get information about the given pokemon
	pokemonName := args[0]
	pokemon, err := conf.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	// 3> Attempt to "catch" the pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	//fmt.Printf("Pokemon id: %d Experience: %d\n", pokemon.Id, pokemon.BaseExperience)
	if attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonName)
		conf.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func attemptCatch(experience int) bool {
	//The more experience, the hard it is to catch.
	max := 250
	attempt := rand.Intn(max)
	//fmt.Printf("Attempt: %d experience: %d\n", attempt, experience)
	return attempt > experience
}
