package main

import (
	"time"

	"github.com/Suppur/pokedexcli/internal/pokecache"
)

func cacheInit() *pokecache.Cache {
	cache := pokecache.NewCache(5 * time.Second)
	return cache
}

func main() {
	cacheInit()
	replInit()

}
