package main

import "fmt"

func main() {
	var myMap = map[string]string{}

	myMap["Todd"] = "Hello"
	myMap["George"] = "Goodbye"

	fmt.Println(myMap)
	fmt.Println(myMap == nil)

	var myMap2 = map[string]string{
		"Todd":"Hello",
		"Jenny":"Goodbye",
	}
	fmt.Println(myMap2)
}
