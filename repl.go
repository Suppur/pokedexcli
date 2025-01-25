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
	Next     string
	Previous string
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
			description: "Displays the name of 20 location areas in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon World",
			callback:    commandMapb,
		},
	}
}

func replInit() {
	reader := bufio.NewScanner(os.Stdin)
	config := &config{
		"https://pokeapi.co/api/v2/location-area/",
		"",
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleaned := cleanInput(reader.Text())
		if len(cleaned) == 0 {
			continue
		}

		cmdString := cleaned[0]
		if cmd, ok := getCommands()[cmdString]; ok {
			err := cmd.callback(config)
			if err != nil {
				fmt.Printf("Unknown command %v", err)
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
