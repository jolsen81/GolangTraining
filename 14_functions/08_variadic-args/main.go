package main

import "fmt"

func main() {
	data := []float64{34, 36, 97, 12, 42, 55}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
