package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"arf"}, true}
	fifi := cat{animal{"meow"}, true}

	specs(fido)
	specs(fifi)
}
