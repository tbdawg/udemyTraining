package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a) // 43
	fmt.Println(&a) // address

	var b *int = &a

	fmt.Println(b)  // address
	fmt.Println(*b) // 43
}
