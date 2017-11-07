package main

import "fmt"

func main() {

	greetings := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(greetings)

	if val, exists := greetings[7]; exists {
		delete(greetings, 7)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(greetings)
}
