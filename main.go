package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("A tiny search engine that gets user input and checks if it present")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("Error reading file", err)
		return
	}

	defer file.Close()

	var search string

	fmt.Print("search: ")
	fmt.Scanln(&search)
	
	scanner := bufio.NewScanner(file)

	line := 1

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(strings.ToLower(text), strings.ToLower(search)) {
			fmt.Println("Found on line: ", line)
		} 
		line++
	}
}

