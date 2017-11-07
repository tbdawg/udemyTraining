package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
	// won't be included in the byteSlice json.Marshal()
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	byteSlice, _ := json.Marshal(p1)
	fmt.Println(byteSlice)
	fmt.Printf("%T \n", byteSlice)
	fmt.Println(string(byteSlice))
}
