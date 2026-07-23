package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	CatchedPokemon   map[string]pokeapi.PokeStats
}

func cleanInput(text string) []string {
	cleanInputs := []string{}
	splitStr := strings.Fields(text)
	for _, word := range splitStr {
		cleanInputs = append(cleanInputs, strings.ToLower(word))
	}
	return cleanInputs
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getAvailibleCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show availible commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "See the pokemon of a selected area",
			callback:    commandExplore,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See the pokemon you have caught",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect catched pokemon stats",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func recieveCommand(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandInput := input[0]
		if len(input) == 1 {
			input = append(input, "")
		}

		availibleCommands := getAvailibleCommands()
		command, ok := availibleCommands[commandInput]
		if ok {
			err := command.callback(cfg, input[1])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
