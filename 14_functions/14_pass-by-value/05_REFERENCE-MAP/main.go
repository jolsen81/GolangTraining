package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Jeremy"])
}

func changeMe(z map[string]int) {
	z["Jeremy"] = 37
}
