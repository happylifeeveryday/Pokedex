package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *config) error {
	url := cfg.Next
	byteData, _, _ := MapCache.Get(url)
	var locationData locationResponse
	json.Unmarshal(byteData, &locationData)
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = locationData.Next
	cfg.Previous = locationData.Previous
	return nil
}

type locationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"json"`
}
