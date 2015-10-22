package main

import "fmt"

func main() {
	fmt.Println("Bitwise AND of 0 and 0 is: ", bitwiseAND(0, 0))
	fmt.Println("Bitwise AND of 0 and 1 is: ", bitwiseAND(0, 1))
	fmt.Println("Bitwise AND of 1 and 0 is: ", bitwiseAND(1, 0))
	fmt.Println("Bitwise AND of 1 and 1 is: ", bitwiseAND(1, 1))
	fmt.Println("Bitwise OR of 0 and 0 is: ", bitwiseOR(0, 0))
	fmt.Println("Bitwise OR of 0 and 1 is: ", bitwiseOR(0, 1))
	fmt.Println("Bitwise OR of 1 and 0 is: ", bitwiseOR(1, 0))
	fmt.Println("Bitwise OR of 1 and 1 is: ", bitwiseOR(1, 1))
}

func bitwiseAND(a, b int) int {
	return a & b
}

func bitwiseOR(a, b int) int {
	return a | b
}
