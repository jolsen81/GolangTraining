package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "007",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Money",
			Last:  "Penny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.Person.First, p1.First)
	fmt.Println(p2.Person.First, p2.First)
}
