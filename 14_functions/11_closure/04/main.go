package main

import "fmt"

func wrapper() func() int {
	// x := 0
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(wrapper()()) // call the function returned by wrapper. closure x value isn't persistent
}
