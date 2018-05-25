package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

func HashBucket(word string, bLength int) int {
	letter := int(word[0])
	bucket := letter % bLength
	return bucket
}
