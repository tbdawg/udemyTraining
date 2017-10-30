package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter a larger number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter a smaller number: ")
	fmt.Scan(&num2)
	fmt.Printf("%v %% %v = %v\n", num1, num2, num1%num2)
}
