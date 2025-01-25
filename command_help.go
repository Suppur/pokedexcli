package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}
