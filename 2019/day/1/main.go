package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bitterpilot/adventofcode/2019/day/1/fuel"
)

func main() {
	file := "./input"
	input := input(file)

	output := fuel.CounterUpper(input)
	fmt.Printf("The required fuel: %0.0f\n", output)
}

func input(file string) []float64 {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	var ret []float64

	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			log.Fatalf("error reading line: %v", err)
		}
		ret = append(ret, i)
	}

	err = s.Err()
	if err != nil {
		log.Fatalf("error scanning file: %v", err)
	}

	return ret
}
