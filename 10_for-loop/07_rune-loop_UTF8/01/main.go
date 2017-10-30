package main

import "fmt"

func main() {
	//ch := 'X'
	//fmt.Printf("rune: %q type: %T value: %d nextRune: %q", ch, ch, ch, ch+1)

	for i := 5000; i <= 5100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	
}
