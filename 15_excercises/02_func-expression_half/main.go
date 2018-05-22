package main

import "fmt"

// func half(z int) (int, bool) {
// 	return z / 2, z%2 == 0
// }

func main() {
	half := func(z int) (int, bool) {
		return z / 2, z%2 == 0
	}
	// half, even := half(4)

	fmt.Println(half(4))
	// fmt.Println(even)
}
