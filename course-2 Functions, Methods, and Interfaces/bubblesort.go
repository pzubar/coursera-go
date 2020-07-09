package main

import (
	"fmt"
	"strconv"
)

func Swap(sli []int, index int) {
	var tmp = sli[index]

	sli[index] = sli[index+1]
	sli[index+1] = tmp
}

func BubbleSort(sli []int) {
	for index := range sli {
		for i := 0; i < len(sli)-index-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func main() {
	var slice = make([]int, 0, 10)
	var userInput string

	fmt.Println("Please, enter up to 10 integers (or press 'X' to sort an array immediately)")
	for len(slice) <= 10 {
		fmt.Scan(&userInput)
		if userInput == "X" {
			break
		} else if s, err := strconv.Atoi(userInput); err == nil {
			slice = append(slice, s)
		}
	}
	BubbleSort(slice)
	fmt.Println(slice)
}
