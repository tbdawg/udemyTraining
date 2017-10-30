package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	//fmt.Println(food) // food is scoped to the if statement
}
