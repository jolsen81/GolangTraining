package main

import "fmt"

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	result := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i

		}
		result <- total
		close(result)
	}()
	return result
}
