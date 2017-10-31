package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(len(students))
	fmt.Println(cap(students))
	fmt.Println(student == nil)
}
