package main

import "fmt"

func Swap(sli []int, index int) {
	var tmp = sli[index]

	sli[index] = sli[index+1]
	sli[index+1] = tmp
}

func BubbleSort(sli []int) {
	for index := range sli {
		for i := index; i < len(sli); i++ {
			if i != len(sli)-1 && sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func main() {
	var test = []int{30, 2, -2, 35, 199, 0}

	BubbleSort(test)
	fmt.Println(test)
}
