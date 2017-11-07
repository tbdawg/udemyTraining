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
	delete(greetings, 7)
	fmt.Println(greetings)
}
