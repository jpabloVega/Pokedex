package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient:  pokeClient,
		CatchedPokemon: make(map[string]pokeapi.PokeStats),
	}
	recieveCommand(cfg)
}
