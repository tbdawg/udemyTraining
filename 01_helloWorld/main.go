package main

import "fmt"

func main() {
	bytesWritten, err := fmt.Println("Hello, World!")

	if err == nil {
		fmt.Println(bytesWritten)
	}

}
