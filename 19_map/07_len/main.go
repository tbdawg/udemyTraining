package main

import "fmt"

func main() {
	greeting := map[string]string{
		"Tim"  : "Good morning.",
		"Jenny": "Bonjour.",
	}

	greeting["Harleen"] = "Howdy"

	fmt.Println(len(greeting))
}
