package main

import (
	"encoding/json"
	"fmt"
)

func commandMapb(cfg *config) error {
	url := cfg.Previous
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
