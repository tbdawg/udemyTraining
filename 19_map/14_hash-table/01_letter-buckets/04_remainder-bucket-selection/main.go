package main

import "fmt"

func main() {
	// outputs the decimal value of the rune (character)
	// then the rune itself and which bucket it goes in
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
