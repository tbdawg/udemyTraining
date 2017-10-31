package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	// xs := []int{}
	var xs []int
	for index, value := range numbers {
		if callback(value) {
			xs = append(xs, index+value)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}
