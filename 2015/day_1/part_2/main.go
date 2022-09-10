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

	// Calculating floor and position, then printing the answer in the terminal.
	floor, position := calculateFloorAndPosition(string(f))

	// Now, given the same instructions, find the position of the first character
	// that causes him to enter the basement (floor -1).
	fmt.Printf("reached floor %v at position %v", floor, position)
}

// CalculateFloorAndPosition calculates floor and returns floor (-1) and position.
// The name of the function is long but it's clear what it does.
func calculateFloorAndPosition(s string) (floor int, position int) {
	// He starts on the ground floor (floor 0)
	// and then follows the instructions one character at a time.
	var f int
	var p int

	for i, v := range s {
		// An opening parenthesis, (, means he should go up one floor,
		// and a closing parenthesis, ), means he should go down one floor.
		if string(v) == "(" {
			f++
		}
		if string(v) == ")" {
			f--
		}

		// What is the position of the character that causes Santa to first enter the basement?
		if f == -1 {
			// Because index starts from 0 increment position to 1.
			p = i + 1
			return f, p
		}
	}

	// Returning calculated floor (should be -1) and position.
	return f, p
}
