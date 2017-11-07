package main

import "fmt"

func main() {
	greeting := map[string]string{}
	greeting["Tim"] = "Good morning."
	greeting["Jenny"] = "Bonjour."

	fmt.Println(greeting["Jenny"])
}
