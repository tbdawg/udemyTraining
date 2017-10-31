package main

import "fmt"

func foo(n ...int) {
	fmt.Println(n)
	/*for _, v := range n {
		fmt.Println(v)
	}*/
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
	foo([]int{9, 8, 7, 6, 5}...)
}
