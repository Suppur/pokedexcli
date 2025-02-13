package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/Suppur/pokedexcli/internal/pokeapi"
)

func commandCatch(c *config, s []string) error {
	if len(s) <= 1 {
		return errors.New("please enter a pokemon name")
	}

	pokeResp, err := c.pokeapiClient.CatchList(s[1])
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %v...\n", pokeResp.Name)
	if (float64(pokeResp.BaseExperience/10) * rand.Float64()) > 5 {
		fmt.Printf("%v was caught!\n", pokeResp.Name)
		if c.pokedex.Caught == nil {
			c.pokedex.Caught = make(map[string]pokeapi.Pokemons)
		}

		c.pokedex.Caught[strings.ToLower(pokeResp.Name)] = pokeResp
		return nil
	}
	//fmt.Printf("pokemon base exp: %v, pokemon chance to be caught: %v\n", pokeResp.BaseExperience, float64(pokeResp.BaseExperience/10)*rand.NormFloat64())
	fmt.Printf("%v escaped!\n", pokeResp.Name)

	return nil
}
