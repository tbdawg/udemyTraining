package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := 'a' // rune
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
	fmt.Printf("%v \n", foo)
	fmt.Printf("%q \n", foo)
	fmt.Printf("%s \n", string(foo))
}
