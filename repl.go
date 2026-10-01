package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
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
