package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bitterpilot/adventofcode/2019/day/2/computer"
)

func main() {
	file := "./input"
	initialState := input(file)

	noun, verb, err := computer.FindSolution(19690720, 99, initialState)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("The Pair:", noun, verb)
}

func input(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	var ret []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		pro := strings.Split(l, ",")

		for _, num := range pro {
			d, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Failed to convert to Int: %v", err)
			}

			ret = append(ret, d)
		}
	}

	err = s.Err()
	if err != nil {
		log.Fatalf("error scanning file: %v", err)
	}

	return ret
}
