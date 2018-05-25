package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	// res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")

	if err != nil {
		log.Fatal(err)
	}

	byteSlice, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", byteSlice)
}
