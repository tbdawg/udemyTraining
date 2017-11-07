package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	// tag, don't add Last when Marshal()'ed
	Last string `json:"-"`
	// tag, change variable identifier Age, to "wisdom score" when Marshal()'ed
	Age int `json:"wisdom score"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
