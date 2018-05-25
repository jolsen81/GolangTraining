package main

import (
	"bufio"
	"fmt"
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

	//Initialize a scanner and scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Create slit to hold counts
	// buckets := make([]int, 12)

	//Loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// n := HashBucket(scanner.Text(), len(buckets))
		// buckets[n]++
	}
	// fmt.Println(buckets)
	// fmt.Println("************************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Printf("%v - %c - %v\n", i, i, buckets[i])
	// }
}
