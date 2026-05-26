package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	location_endpoint = "https://pokeapi.co/api/v2/location-area/"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

type locationResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
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
			callback:    commandMapB,
		},
	}
}

func getCallback(command string) func(*config) error {
	if val, ok := commands[command]; ok {
		return val.callback
	}
	return func(*config) error {
		fmt.Println("Unknown command")
		return nil
	}
}

//handlers

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}

func commandMap(config *config) error {
	if len(config.next) == 0 {
		fmt.Println("you're on the first page")
		return mapCommand(config, location_endpoint)
	}
	return mapCommand(config, config.next)

}

func commandMapB(config *config) error {
	if len(config.previous) == 0 {
		fmt.Println("you're on the first page")
		return mapCommand(config, location_endpoint)
	}
	return mapCommand(config, config.previous)

}

func mapCommand(config *config, endpoint string) error {
	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from API", resp.StatusCode)
	}

	var response locationResp
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if len(response.Results) == 0 {
		fmt.Println("No locations found.")
		return nil
	}

	config.next = response.Next
	config.previous = response.Previous

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}
