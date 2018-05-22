package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello, world!"
	}
}

func main() {
	greeting := makeGreeter()
	fmt.Println(greeting())
	fmt.Printf("%T\n", greeting)
}
