package main

import (
	"fmt"
	"log"
)

func greatestNum(n ...int) int {
	if len(n) < 1 {
		fmt.Println("There are no values to evaluate")
		log.Fatal()
	}
	val := n[0]

	for _, temp := range n {
		if temp > val {
			val = temp
		}
	}
	return val
}

func main() {
	greatest := greatestNum([]int{}...)
	fmt.Println(greatest)
}
