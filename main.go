package main

import (
	"time"

	"github.com/Suppur/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	conf := &config{
		pokeapiClient: pokeClient,
	}

	replInit(conf)

}
