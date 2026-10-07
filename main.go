package main

import (
	"time"

	"github.com/Ehs-Colin/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	conf := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(conf)
}
