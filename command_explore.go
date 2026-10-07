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
	areaDetails, err := conf.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range areaDetails.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

// func commandMapForward(conf *config) error {
// 	mapArea, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
// 	if err != nil {
// 		return err
// 	}
// 	conf.nextLocationsURL = mapArea.Next
// 	conf.prevLocationsURL = mapArea.Previous
// 	for _, area := range mapArea.Results {
// 		fmt.Println(area.Name)
// 	}
// 	return nil
// }
