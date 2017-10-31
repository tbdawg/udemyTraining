package main

import "fmt"

func greatest(n ...int) int {
	var g int
	for _, v := range n {
		if v > g {
			g = v
		}
	}
	return g
}

func main() {
	fmt.Println(greatest(42, 64, 24, 633, 34, 75))
}
