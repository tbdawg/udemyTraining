package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(len(students))
	fmt.Println(cap(students))
	fmt.Println(student == nil)
}
