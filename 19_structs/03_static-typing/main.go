package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 37
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	yourAge = 44
	fmt.Printf("%T %v\n", yourAge, yourAge)

	//this doesn't work
	// fmt.Println(myAge + yourAge)

	//this does work
	fmt.Println(int(myAge) + yourAge)
}
