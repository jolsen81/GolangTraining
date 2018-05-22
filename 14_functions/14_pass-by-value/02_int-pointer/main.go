package main

import "fmt"

func main() {
	age := 37
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)

	*z = 27

	fmt.Println(z)
	fmt.Println(*z)
}
