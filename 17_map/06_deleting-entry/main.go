package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning.",
		"Jenny": "Bonjour",
	}
	fmt.Println(myGreeting)
	delete(myGreeting, "Tim")

	fmt.Println(myGreeting)
}
