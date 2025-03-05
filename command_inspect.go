package main

import "fmt"

func commandInspect(cfg *config) error {
	pokemon, exist := Pokedex[cfg.Parameter]
	if exist {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		fmt.Printf("  -hp: %v\n", pokemon.Stats[0].BaseStat)
		fmt.Printf("  -attack: %v\n", pokemon.Stats[1].BaseStat)
		fmt.Printf("  -defense: %v\n", pokemon.Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %v\n", pokemon.Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %v\n", pokemon.Stats[4].BaseStat)
		fmt.Printf("  -speed: %v\n", pokemon.Stats[5].BaseStat)
		fmt.Printf("Types:\n")
		fmt.Printf("  - %v\n", pokemon.Types[0].Type.Name)
		fmt.Printf("  - %v\n", pokemon.Types[1].Type.Name)
	} else {

	}
	return nil
}
