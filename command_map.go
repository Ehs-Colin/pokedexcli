package main

import (
	"fmt"
)

type MapArea struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}

func commandMapForward(conf *config) error {
	mapArea, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}
	conf.nextLocationsURL = mapArea.Next
	conf.prevLocationsURL = mapArea.Previous
	for _, area := range mapArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.prevLocationsURL == nil {
		return fmt.Errorf("You're on the first page")
	}
	mapArea, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}
	conf.nextLocationsURL = mapArea.Next
	conf.prevLocationsURL = mapArea.Previous
	for _, area := range mapArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}
