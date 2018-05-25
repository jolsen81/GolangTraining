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

	if err != nil {
		log.Fatal(err)
	}

	//Initialize a scanner and scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	fmt.Println(res)

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Create slit to hold counts
	buckets := make([]int, 200)

	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	fmt.Println("************************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Printf("%v - %c - %v\n", i, i, buckets[i])
	// }
}

func HashBucket(w string) int {
	return int(w[0])
}
