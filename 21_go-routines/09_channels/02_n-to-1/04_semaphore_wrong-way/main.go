package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// the below three lines of code will block progress and c will not have anything
	//taken off the channel, freezing the program. These lines of code should be
	//placed within a block go function to keep program from freezing
	<-done
	<-done
	close(c)

	for n := range c {
		fmt.Println(n)
	}

}
