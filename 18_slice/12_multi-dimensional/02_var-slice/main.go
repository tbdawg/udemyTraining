package main

import "fmt"

func main() {
	var student []string
	var students [][]string

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(len(students))
	fmt.Println(cap(students))
	fmt.Println(student == nil)
}
