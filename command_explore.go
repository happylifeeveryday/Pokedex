package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(cfg *config) error {
	prefix := cfg.Parameter
	url := "https://pokeapi.co/api/v2/location-area/" + prefix
	fmt.Printf("Exploring %v...\n", prefix)
	byteData, _, err := Cache.Get(url)
	if err != nil {
		return err
	}
	var locationDetailsData locationDetails
	err = json.Unmarshal(byteData, &locationDetailsData)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonInfo := range locationDetailsData.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemonInfo.Pokemon.Name)
	}
	return nil
}

type locationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
