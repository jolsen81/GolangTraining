package main

import "fmt"

func zero(x *int) {
	fmt.Println(*x)
	fmt.Println(x)
	fmt.Println()
	*x = 0
}

func main() {
	x := 5

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println()

	zero(&x)
	fmt.Println(x)
}
