package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%3d \t %08b \t %#X \t %q \n", i, i, i, i)
	}
}
