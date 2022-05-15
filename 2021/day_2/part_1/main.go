package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	var forward int
	var aim int
	var depth int
	for s.Scan() {
		f := strings.HasPrefix(s.Text(), "forward ")
		up := strings.HasPrefix(s.Text(), "up ")
		down := strings.HasPrefix(s.Text(), "down ")
		if f {
			a := strings.TrimPrefix(s.Text(), "forward ")
			x, err := strconv.Atoi(a)
			if err != nil {
				log.Fatal(err)
			}
			forward = forward + x
			depth = depth + aim*x
		}

		if up {
			b := strings.TrimPrefix(s.Text(), "up ")
			y, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}
			// depth = depth - y
			aim = aim - y
		}

		if down {
			c := strings.TrimPrefix(s.Text(), "down ")
			z, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			// depth = depth + z
			aim = aim + z
		}

	}
	fmt.Println(forward)
	fmt.Println(depth * forward)
	fmt.Println(aim)
}
