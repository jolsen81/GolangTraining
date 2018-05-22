package main

import "fmt"

func main() {
	greet("Jeremy", "Olsen")
}

func greet(fname string, lname string) {
	fmt.Println("Hello,", fname, lname)
}
