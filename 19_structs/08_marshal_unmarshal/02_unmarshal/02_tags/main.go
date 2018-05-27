package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	byteSlice := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(byteSlice, &p1)

	fmt.Println("--------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}
