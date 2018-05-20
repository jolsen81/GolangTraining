package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  //prints value of a
	fmt.Println(&a) //prints memory address of a

	var b *int = &a //b points to memory address of a

	fmt.Println(b)  //prints memory address b points to = &a
	fmt.Println(*b) //dereferences the address b points to (&a) and prints the value stored in that address (a)

	*b = 25 //dereferences the address b points to (&a) and changes the value at that address

	fmt.Println(a) //prints new value for a
}
