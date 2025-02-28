package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	url := cfg.Next
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var locationData locationResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationData)
	if err != nil {
		return err
	}
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
