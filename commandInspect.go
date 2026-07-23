package main

import (
	"fmt"
)

func commandInspect(cfg *config, param1 string) error {
	if pokemon, ok := cfg.CatchedPokemon[param1]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, tpe := range pokemon.Types {
			fmt.Printf("  -%s\n", tpe.Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
}
