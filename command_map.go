package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	var loc locations

	err = json.Unmarshal([]byte(body), &loc)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}
	if loc.Previous != nil {
		c.Previous = *loc.Previous
	} else {
		c.Previous = ""
	}
	c.Next = *loc.Next

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(c.Previous)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	var loc locations

	err = json.Unmarshal([]byte(body), &loc)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	c.Next = *loc.Next
	if loc.Previous != nil {
		c.Previous = *loc.Previous
	} else {
		c.Previous = ""
	}

	return nil
}
