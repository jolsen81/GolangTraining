package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			// fmt.Println(<-c)
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
			// c <- 8
		}
	}()

	time.Sleep(time.Second)
}
