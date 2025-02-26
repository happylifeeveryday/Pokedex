package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Printf("Hello, World!")
}

func cleanInput(text string) []string{
	// Trim leading/trailing whitespace and convert the text to lowercase
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// Split the text by whitespace into a slice of words
	words := strings.Fields(text)

	return words
}
