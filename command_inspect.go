package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(c *config, s []string) error {
	if len(s) <= 1 {
		return errors.New("please enter a pokemon name")
	}

	pokemon := strings.ToLower(s[1])
	poke, ok := c.pokedex.Caught[pokemon]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", poke.Name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Printf("Weight: %v\n", poke.Weight)
	fmt.Println("Stats: ")
	for _, stat := range poke.Stats {
		fmt.Printf("	- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, typ := range poke.Types {
		fmt.Printf("	- %v\n", typ.Type.Name)
	}
	return nil
}
