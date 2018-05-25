package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	temp := myGreeting
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
	fmt.Println(temp)
}
