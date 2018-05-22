package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Enter a number: ")
	fmt.Scan(&x)

	fmt.Print("Enter another number: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println(x % y)
	} else {
		fmt.Println(y % x)
	}
}
