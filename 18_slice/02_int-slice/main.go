package main

import "fmt"

func main() {
	sl := []int{1, 3, 5, 7, 9, 11,}

	for i, value := range sl {
		fmt.Println(i, "-", value)
	}
}
