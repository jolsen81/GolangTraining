package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Jeremy")
	var c = Contact{"Good to see you", "Time"}
	SwitchOnType(c)
	SwitchOnType(true)
}
