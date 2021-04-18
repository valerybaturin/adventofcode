package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Reading file to variable.
	f, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Calculating floor and printing answer in terminal.
	fmt.Println(CalculateFloor(string(f)))
}

func CalculateFloor(s string) int {
	// He starts on the ground floor (floor 0)
	// and then follows the instructions one character at a time.
	var floor int

	for _, v := range s {
		// An opening parenthesis, (, means he should go up one floor,
		// and a closing parenthesis, ), means he should go down one floor.
		if string(v) == "(" {
			floor++
		}
		if string(v) == ")" {
			floor--
		}
	}

	// Returning calculated floor
	return floor
}
