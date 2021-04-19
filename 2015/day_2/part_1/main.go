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

// Area calculating the required wrapping paper for each gift
func (b *Box) Area() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

// SmallArea the area of the smallest side
func (b *Box) SmallArea() int {
	s := []int{b.l, b.w, b.h}
	sort.Ints(s)
	return s[0] * s[1]
}

// ParseDimensions parsing input and returning Box type
func ParseDimensions(s string) *Box {
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
		box := ParseDimensions(s.Text())
		// Calculate wrapping paper plus little extra paper for each present
		area := box.Area() + box.SmallArea()
		// Sum to previous value
		answer = answer + area
	}

	fmt.Println(answer)
}
