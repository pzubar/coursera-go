package main

import (
	"fmt"
	"strings"
)

func findian(userInput string) {
	if strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

func main() {
	var userInput string
	fmt.Println("Please, enter a string")
	fmt.Scan(&userInput)

	findian(userInput)
}
