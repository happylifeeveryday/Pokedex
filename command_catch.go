package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", cfg.Parameter)
	url := "https://pokeapi.co/api/v2/pokemon/" + cfg.Parameter
	byteData, _, _ := Cache.Get(url)
	var pokemon Pokemon
	err := json.Unmarshal(byteData, &pokemon)
	if err != nil {
		return nil
	}
	//construct rand
	rand := rand.New(rand.NewSource(99))
	prob := float64(rand.Intn(100)) - float64(pokemon.BaseExperience)/5
	if prob > 0 {
		Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%v was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}
