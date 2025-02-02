package main

import (
	"errors"
	"fmt"
)

func commandMapf(c *config, s []string) error {
	locationsRes, err := c.pokeapiClient.ListLocations(c.nextURL)
	if err != nil {
		return err
	}

	c.nextURL = locationsRes.Next
	c.previousURL = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, s []string) error {
	if c.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationsRes, err := c.pokeapiClient.ListLocations(c.previousURL)
	if err != nil {
		return err
	}

	c.nextURL = locationsRes.Next
	c.previousURL = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
