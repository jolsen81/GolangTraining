package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 37
	fmt.Printf("%T %v\n", myAge, myAge)
}
