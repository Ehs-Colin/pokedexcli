package main

import (
	"time"

	"github.com/Ehs-Colin/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(conf)
}
