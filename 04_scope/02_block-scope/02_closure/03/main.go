package main

import "fmt"

func main() {
	x := 0

	//anonymous function
	increment := func() int {
		x++
		return x

	}

	fmt.Println(increment())
	fmt.Println(increment())
}
