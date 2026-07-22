package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, param1 string) error {
	if param1 == "" {
		err := errors.New("Correct syntaxis: explore <location name>")
		return err
	}

	pokemonRes, err := cfg.pokeapiClient.ListEncounters(param1)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", param1)
	fmt.Println("Found pokemon:")
	for _, encounter := range pokemonRes.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
	return nil
}
