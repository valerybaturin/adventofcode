package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Box contains present dimensions
type Box struct {
	l, w, h int
}

// NewBox is a constructor returns new Box
func NewBox(l, w, h int) *Box {
	return &Box{l: l, w: w, h: h}
}

// Ribbon calculates ribbon to wrap the present
func (b *Box) Ribbon() int {
	s := []int{b.l, b.w, b.h}
	sort.Ints(s)
	return s[0] + s[0] + s[1] + s[1]
}

// Bow calculates ribbon for the bow
func (b *Box) Bow() int {
	return b.l * b.w * b.h
}

// ParseDimensions parsing input and returning Box type
func parseDimensions(s string) *Box {
	d := strings.Split(s, "x")
	di := make([]int, 3)

	for i, v := range d {
		di[i], _ = strconv.Atoi(v)
	}

	return NewBox(di[0], di[1], di[2])
}

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

	// This variable holds our answer.
	var answer int

	// Scanning file content
	s := bufio.NewScanner(file)
	for s.Scan() {
		// Get box dimensions from string
		r := parseDimensions(s.Text())
		// Calculate wrapping ribbon
		length := r.Ribbon() + r.Bow()
		// Sum to previous value
		answer = answer + length
	}

	fmt.Println(answer)
}
