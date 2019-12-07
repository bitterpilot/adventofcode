package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bitterpilot/adventofcode/2019/day/3/wires"
)

func main() {
	file := "./input"
	initialState := input(file)

	solution := wires.Solve(initialState[0], initialState[1])

	fmt.Printf("The Solution: %d\n", solution)
}

func input(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	var ret [][]string

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		pro := strings.Split(l, ",")

		ret = append(ret, pro)

	}

	err = s.Err()
	if err != nil {
		log.Fatalf("error scanning file: %v", err)
	}

	return ret
}
