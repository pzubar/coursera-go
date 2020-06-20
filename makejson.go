package main

import (
	"encoding/json"
	"fmt"
)

func askUser(prompt string, userInput string) string {
	fmt.Println(prompt)
	_, err := fmt.Scan(&userInput)
	if err == nil {
		return userInput
	} else {
		return ""
	}
}

func main() {
	var personMap map[string]string
	var userInput string

	personMap = make(map[string]string)
	personMap["name"] = askUser("Enter your name, please:", userInput)
	personMap["address"] = askUser("Enter your address, please:", userInput)
	jsonString, err := json.Marshal(personMap)
	if err == nil {
		fmt.Println(string(jsonString))
	} else {
		fmt.Println("An error occurred!")
	}
}
