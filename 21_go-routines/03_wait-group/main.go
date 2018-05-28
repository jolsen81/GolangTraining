package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		println("Bar:", i)
	}

	wg.Done()
}
