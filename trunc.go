package main

import "fmt"

func trunc() {
	var userInput float64

	fmt.Println("Please, enter a floating point number")
	fmt.Scan(&userInput)
	fmt.Println(int(userInput))
}

func main() {
	trunc()
}
