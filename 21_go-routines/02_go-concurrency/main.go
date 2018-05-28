package main

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 1000; i++ {
		println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 1000; i++ {
		println("Bar:", i)
	}
}
