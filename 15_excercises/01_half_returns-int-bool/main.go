package main

import "fmt"

func half(z int) (int, bool) {
	return z / 2, z%2 == 0
}

func main() {
	val, even := half(4)

	fmt.Println(val)
	fmt.Println(even)
}
