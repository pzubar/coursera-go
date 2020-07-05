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

func getNamesFromFile(file *os.File) []Name {
	var names []Name
	var scanner = bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		for scanner.Scan() {
			words := strings.Split(scanner.Text(), " ")

			if len(words) == 2 {
				names = append(names, Name{words[0], words[1]})
			} else {
				fmt.Println("Wrong format of string!")
			}
		}
	}

	return names
}

func main() {
	var names []Name
	var filename string

	fmt.Println("Please, provide a filename")
	_, err := fmt.Scan(&filename)
	file, readError := os.Open(filename)

	if err != nil || readError != nil {
		fmt.Println("An error occurred!")
		return
	}

	names = getNamesFromFile(file)
	for _, fullName := range names {
		fmt.Printf("%s %s\n", fullName.fname, fullName.lname)
	}

	closingError := file.Close()
	if closingError != nil {
		fmt.Println("An error occurred while closing the file")
	}
}
