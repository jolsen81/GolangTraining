package main

import (
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // no longer necessary to run this
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		println("Bar:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
