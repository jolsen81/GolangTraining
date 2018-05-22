package main

import "fmt"

func main() {
	name := "Jeremy"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)
}

func changeMe(s string) {
	fmt.Println(s)
	s = "Rocky"
	fmt.Println(s)
}
