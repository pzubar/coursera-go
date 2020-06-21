package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var names []Name
	file, err := os.Open("test")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		fmt.Println(words, len(words))
		names = append(names,
			Name{
				words[0], words[1],
			})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)
	file.Close()
}
