package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	in := make(chan string)

	go workerProcess(in)
	go workerProcess(in)
	go workerProcess(in)
	go publisher(in)
	go publisher(in)
	go publisher(in)
	go publisher(in)
	time.Sleep(1 * time.Millisecond)
}

//publisher pushes data into a channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: wating for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisId, input)
	}
}
