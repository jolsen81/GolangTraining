package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
	// notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	byteSlice, _ := json.Marshal(p1)
	// fmt.Println(byteSlice)
	// fmt.Printf("%T\n", byteSlice)
	fmt.Println(string(byteSlice))
}
