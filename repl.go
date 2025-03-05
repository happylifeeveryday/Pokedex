package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/happylifeeveryday/Pokedex/pokecache"
)

var Cache = pokecache.NewCache(600 * time.Second)
var Pokedex = make(map[string]Pokemon)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "https://pokeapi.co/api/v2/location-area",
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if commandName == "explore" || commandName == "catch" {
			if len(words) == 2 {
				parameter := words[1]
				cfg.Parameter = parameter
			} else {
				fmt.Println("Wrong parameter")
				continue
			}
		}

		command, exist := getCommands()[commandName]
		if exist {
			command.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next      string
	Previous  string
	Parameter string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
	}
}
