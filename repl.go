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
	callback    func() error
}

var commandMap map[string]cliCommand

func init() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
	}
}

func replInit() {
	reader := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleaned := cleanInput(reader.Text())
		if len(cleaned) == 0 {
			continue
		}
		cmdString := cleaned[0]
		if cmd, ok := commandMap[cmdString]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Errorf("Unknown command %v", err)
				continue
			}
		}else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commandMap {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}

func cleanInput(text string) []string {
	fmtString := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return fmtString
}
