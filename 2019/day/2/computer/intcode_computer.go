package computer

import "log"

// Intcode is a program that reads a array of integers to perform operations.
//
// To run start buy looking at the first integer(position 0) this is an optcode.
// The opcode indicates what to do; 1 adds, 2 multiplies and 99 halts the program.
// after reading the opt code the next 2 positions contain references to the
// operands and the third is a reference to where the result is stored.
// For example, if your Intcode computer encounters 1,10,20,30,
// it should read the values at positions 10 and 20, add those values,
// and then overwrite the value at position 30 with their sum.
func Intcode(program []int) []int {
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
			return program
		default:
			log.Fatalf("Unknown opt code %d at position %d", program[opt], opt)
		}
	}

	return program
}
