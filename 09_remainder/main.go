package main

import (
	"fmt"
)

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 1; i <= 50; i++ {
		if i % 2 == 1 {
			fmt.Println(i, "Odd")
		} else {
			fmt.Println(i, "Even")
		}
	}
}
