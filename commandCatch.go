package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, param1 string) error {
	if param1 == "" {
		err := errors.New("Correct syntaxis: catch <pokemon name>")
		return err
	}

	pokemonRes, err := cfg.pokeapiClient.ListStats(param1)
	if err != nil {
		return err
	}

	name := pokemonRes.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	catchChance := rand.IntN(800-18) + 18

	if pokemonRes.BaseExperience > catchChance {
		fmt.Printf("%s escaped!", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.CatchedPokemon[name] = pokemonRes

	return nil

}
