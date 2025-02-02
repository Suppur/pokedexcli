package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, s []string) error {
	if s == nil {
		return errors.New("please enter a location name")
	}

	explRes, err := c.pokeapiClient.ExploreList(s[1])
	if err != nil {
		return err
	}
	fmt.Printf("\nExploring %v...\n", s[1])
	fmt.Printf("Found Pokemon:\n")
	for _, pokemon := range explRes.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
