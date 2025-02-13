package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, s []string) error {
	if len(c.pokedex.Caught) == 0 {
		return errors.New("no pokemon have been caught")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex.Caught {
		fmt.Printf("	- %v\n", pokemon.Name)
	}

	return nil
}
