package main

import (
	"runtime"
	"sync"
	// "time"
)

var wg sync.WaitGroup
// var mutex sync.Mutex
var counter = 0

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // no longer necessary to run this
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	println(counter)
}

func foo() {
	for i := 0; i < 100000; i++ {
    // mutex.Lock()
		counter++
    // mutex.Unlock()
		// println("Foo:", i)
		// time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100000; i++ {
    // mutex.Lock()
		counter++
    // mutex.Unlock()
		// println("Bar:", i)
		// time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
