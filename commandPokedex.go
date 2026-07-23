package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, param1 string) error {
	if len(cfg.CatchedPokemon) == 0 {
		err := errors.New("You have no pokemon")
		return err
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CatchedPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
