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
	// fmt.Println(res)

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Create slit to hold counts
	buckets := make([][]string, 12)
	//create slices to hold words
	fmt.Println("1")
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	//Loop over the words
	fmt.Println("2")
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		word := scanner.Text()
		n := HashBucket(word, len(buckets))
		buckets[n] = append(buckets[n], word)
	}
	// fmt.Println("3")
	//Print len of each bucket
	for i := 0; i < len(buckets); i++ {
		fmt.Println(i, "-", len(buckets[i]))
	}

	fmt.Println(buckets[6])
}

func HashBucket(w string, nBuckets int) int {
	var sum int
	for _, v := range w {
		sum += int(v)
	}
	return sum % nBuckets
}
