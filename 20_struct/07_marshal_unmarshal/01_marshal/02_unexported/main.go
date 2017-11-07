package main

import (
	"fmt"
	"encoding/json"
)

// Person won't be included in the byteSlice json.Marshal()
type Person struct {
	first string
	last string
	age int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	byteSlice, _ := json.Marshal(p1) // stringify
	fmt.Println(string(byteSlice))
}
