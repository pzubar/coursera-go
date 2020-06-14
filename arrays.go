package main

import "fmt"

func main() {

	idMap := map[string]int{
		"one":  2,
		"otto": 4,
	}

	idMap["lol"] = 2

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
