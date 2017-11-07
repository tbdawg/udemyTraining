package main

import "fmt"

func main() {
	greeting := map[string]string{
		"zero" : "Good morning!",
		"one"  : "Bonjour!",
		"two"  : "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(greeting)
	delete(greeting, "two")
	fmt.Println(greeting)
}
