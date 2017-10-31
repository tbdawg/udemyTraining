package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)

	student[0] = "Todd" // student already has 35 slots available so this will work

	fmt.Println(student)
	fmt.Println(students)
}
