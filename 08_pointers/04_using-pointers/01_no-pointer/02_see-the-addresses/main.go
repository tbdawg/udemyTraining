package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Printf("zero: %p\n", &z)
	fmt.Println("zero:", &z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("main: %p\n", &x)
	fmt.Println("main:", &x)
	zero(x)
	fmt.Println(x) // x is still 5
}
