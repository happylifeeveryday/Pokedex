package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		formatted := cleanInput(input)[0]
		fmt.Printf("Your command was: %v\n", formatted)
	}
}

func cleanInput(text string) []string {
	// Trim leading/trailing whitespace and convert the text to lowercase
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// Split the text by whitespace into a slice of words
	words := strings.Fields(text)

	return words
}
