package computer

// Intcode is a program that reads a array of integers to perform operations.
//
// To run start buy looking at the first integer(position 0) this is an optcode.
// The opcode indicates what to do; 1 adds, 2 multiplies and 99 halts the program.
// after reading the opt code the next 2 positions contain references to the 
// operands and the third is a reference to where the result is stored.
// For example, if your Intcode computer encounters 1,10,20,30,
// it should read the values at positions 10 and 20, add those values,
// and then overwrite the value at position 30 with their sum.
func Intcode([]int) []int {
	return nil
}
