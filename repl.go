package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Suppur/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
	pokedex       *pokeapi.Pokedex
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
		"mapf": {
			name:        "mapf",
			description: "Displays the name of 20 location areas in the Pokemon World",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon World",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the pokmon in a given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a previously caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all your captured pokemon",
			callback:    commandPokedex,
		},
	}
}

func replInit(conf *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleaned := cleanInput(reader.Text())
		if len(cleaned) == 0 {
			continue
		}

		cmdString := cleaned[0]
		if cmd, ok := getCommands()[cmdString]; ok {
			err := cmd.callback(conf, cleaned)
			if err != nil {
				fmt.Printf("Unknown command %v\n", err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	fmtString := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return fmtString
}
