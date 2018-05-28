package main

import "fmt"

func main() {

	go func() {
		c := factorial(i % 10)
		for n := range c {
			fmt.Println("Channel:", i, i%8, "- Squared:", n)
		}
	}()

}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= n
		}
		out <- total
		close(out)
	}()
	return out
}
