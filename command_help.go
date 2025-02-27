package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	return nil
}
