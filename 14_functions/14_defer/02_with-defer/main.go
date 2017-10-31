package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Print("world")
}

func ex() {
	fmt.Println("!")
}

func main() {
	defer ex()
	defer world()
	hello()
}
