package main

import "fmt"

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range Pokedex {
		fmt.Printf("- %v\n", key)
	}
	return nil
}
