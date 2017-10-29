package main

import "fmt"

func main() {
	// %x lowercase hexadecimal - %X uppercase hexadecimal
	fmt.Printf("%X\n", 42)
	// # adds 0x/0X prefix for hex and leading 0 for Octal
	fmt.Printf("%#X\n", 42)
}
