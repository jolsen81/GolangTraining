package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
