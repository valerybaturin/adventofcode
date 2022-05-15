package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file.
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close file before main function exits.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// Create new buffer for file.
	s := bufio.NewScanner(file)

	// This variable will hold numbers from input file.
	var depths []int

	// Scan file.
	for s.Scan() {
		// Our input contains string, so we need to convert it to int.
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		// Add number to our variable
		depths = append(depths, v)
	}

	// This will be our answer.
	answer := countDepthIncrease(abc(depths))

	// Print our answer, well done!
	fmt.Println(answer)
}

func abc(depths []int) (sums []int) {
	for i := 0; i < len(depths)-2; i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		sums = append(sums, sum)
	}
	return sums
}

func countDepthIncrease(depths []int) (increased int) {
	for i := range depths {
		if i == 0 {
			continue
		}
		if depths[i] > depths[i-1] {
			increased++
		}
	}

	return increased
}
