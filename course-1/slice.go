package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var userInput string
	slice := make([]int, 0, 3)

	for true {
		fmt.Println("Please, enter an integer")
		fmt.Scan(&userInput)
		if userInput == "X" {
			break
		} else if s, err := strconv.Atoi(userInput); err == nil {
			slice = append(slice, s)
			sort.Ints(slice)
			fmt.Println(slice)
		}
	}
}
