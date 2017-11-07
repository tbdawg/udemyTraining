package main

import "fmt"

func main() {
	greeting := map[string]string{
		"Tim"  : "Good morning.",
		"Jenny": "Bonjour.",
	}

	fmt.Println(greeting["Jenny"])
}
