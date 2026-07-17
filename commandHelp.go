package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, opc := range getAvailibleCommands() {
		fmt.Printf("%v: %v\n", opc.name, opc.description)
	}
	return nil
}
