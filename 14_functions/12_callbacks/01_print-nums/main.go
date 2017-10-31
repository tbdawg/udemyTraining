package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers { // range returns index, value at index or key, value
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
