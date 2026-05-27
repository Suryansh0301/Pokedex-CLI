package main

import (
	"fmt"
	"os"

	"github.com/Suryansh0301/pokedexcli/internal/pokeapi"
)

const (
	location_endpoint = "https://pokeapi.co/api/v2/location-area/"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in pokemon world",
			callback:    commandMapB,
		},
	}
}

func getCallback(command string) func(*pokeapi.Config) error {
	if val, ok := commands[command]; ok {
		return val.callback
	}
	return func(*pokeapi.Config) error {
		fmt.Println("Unknown command")
		return nil
	}
}

//handlers

func commandExit(_ *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}

func commandMap(config *pokeapi.Config) error {
	endpoint := config.Next
	if len(endpoint) == 0 {
		fmt.Println("you're on the first page")
		endpoint = location_endpoint
	}

	client := pokeapi.NewClient()
	resp, err := client.LocationAreas(config, endpoint)
	if err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	return nil

}

func commandMapB(config *pokeapi.Config) error {
	endpoint := config.Previous
	if len(endpoint) == 0 {
		fmt.Println("you're on the first page")
		endpoint = location_endpoint
	}

	client := pokeapi.NewClient()
	resp, err := client.LocationAreas(config, endpoint)
	if err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	return nil

}
