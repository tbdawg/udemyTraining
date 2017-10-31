package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

	// get first element from mySlice and append all elements from first element on, of myOtherSlice
	// mySlice = append(mySlice[:1], myOtherSlice[1:]...)
	// fmt.Println(mySlice)
}
