package main

import "fmt"

const a = 7

const (
	A = iota
	B
	C

)

const (
	_ = iota
	D = iota + C
	E = iota + D
	F = iota + D
)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {
	const b = 4
	fmt.Printf("a: %v b: %v\n", a, b)
	fmt.Printf("A: %v B: %v C: %v\n", A, B, C)
	fmt.Printf("D: %v E: %v F: %v\n", D, E, F)
	fmt.Printf("%b \t %d\n", KB, KB)
	fmt.Printf("%b \t %d\n", MB, MB)
}
