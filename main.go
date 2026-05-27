package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Suryansh0301/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	globalConfig := &pokeapi.Config{}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		callbackFunc := getCallback(scanner.Text())
		err := callbackFunc(globalConfig)
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
