package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ehs-Colin/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of the next 20 location areas in the Pokemon world",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in teh Pokemon world",
			callback:    commandMapBack,
		},
	}
}

func startRepl(conf *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		if err := reader.Err(); err != nil {
			fmt.Printf("IO Error: %s", err.Error())
		}
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		if command, exists := conf.commands[commandName]; exists {
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			//Error, unknown command
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	split_strings := strings.Fields(strings.ToLower(text))
	return split_strings
}
