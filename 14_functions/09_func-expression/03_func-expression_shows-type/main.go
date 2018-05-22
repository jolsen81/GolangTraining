package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello, world!")
	}
	fmt.Printf("%T\n", greeting)
	greeting()
}
