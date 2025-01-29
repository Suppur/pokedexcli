package main

import (
	"time"

	"github.com/Suppur/pokedexcli/internal/pokeapi"
	"github.com/Suppur/pokedexcli/internal/pokecache"
)

func cacheInit() *pokecache.Cache {
	cache := pokecache.NewCache(5 * time.Second)
	return cache
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		pokeapiClient: pokeClient,
	}

	cacheInit()
	replInit(conf)

}
