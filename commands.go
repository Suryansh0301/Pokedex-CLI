package main

import (
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
}

var commands map[string]cliCommand

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
			callback:    commandMap,
		},
	}
}

func getCallback(command string) func() error {
	if val, ok := commands[command]; ok {
		return val.callback
	}
	return func() error {
		fmt.Println("Unknown command")
		return nil
	}
}

//handlers

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}

func commandMap(*config) error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("encountered error while fetching the data: %s", err.Error())
	}
	respBytes := make([]byte, 200)
	resp.Body.Read(respBytes)
	return nil
}
