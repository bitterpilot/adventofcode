package computer

import "fmt"

// Intcode is a program that reads a array of integers to perform operations.
//
// To run start buy looking at the first integer(position 0) this is an optcode.
// The opcode indicates what to do; 1 adds, 2 multiplies and 99 halts the program.
// after reading the opt code the next 2 positions contain references to the
// operands and the third is a reference to where the result is stored.
// For example, if your Intcode computer encounters 1,10,20,30,
// it should read the values at positions 10 and 20, add those values,
// and then overwrite the value at position 30 with their sum.
func Intcode(program []int) ([]int, error) {
	const (
		add      int = 1
		multiply int = 2
		halt     int = 99
	)

	var (
		val1Ptr int
		val2Ptr int
		destPtr int
	)

	step := 4 // The number of postions until the next opt code
	opt := 0
	for true {
		switch program[opt] {
		case add:
			val1Ptr = program[opt+1]
			val2Ptr = program[opt+2]
			destPtr = program[opt+3]
			result := program[val1Ptr] + program[val2Ptr]

			program[destPtr] = result
			opt = opt + step
		case multiply:
			val1Ptr = program[opt+1]
			val2Ptr = program[opt+2]
			destPtr = program[opt+3]
			result := program[val1Ptr] * program[val2Ptr]
			program[destPtr] = result
			opt = opt + step
		case halt:
			return program, nil
		default:
			return program, fmt.Errorf("Unknown opt code %d at position %d", program[opt], opt)
		}
	}

	return program, nil
}

// FindSolution searches for the target value by iterating through ALL possible values.
func FindSolution(target, max int, initialState []int) (int, int, error) {
	count := 0
	verb := 0
	for verb <= max {
		noun := 0
		for noun <= max {
			// reset state
			program := append([]int(nil), initialState...)

			// set test pair
			program[1] = verb
			program[2] = noun

			// Run
			out, err := Intcode(program)
			if err != nil {
				return -1, -1, err
			}

			// check output
			if out[0] == target {
				return verb, noun, nil
			}
			count++
			noun++
		}
		verb++
	}
	fmt.Println(count)
	return -1, -1, fmt.Errorf("solution not found")
}

// PrintProgram is a quick and dirty printer.
// It prints the program to console and only works for programs that have a step of 4.
func PrintProgram(program []int) {
	opt := 0
	for opt < len(program) {
		code := program[opt : opt+4]
		fmt.Printf("line: %3d, %v\n", opt, code)
		opt = opt + 4
	}
}
