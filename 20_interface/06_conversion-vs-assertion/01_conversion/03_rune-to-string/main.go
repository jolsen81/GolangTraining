package main

func main() {
	var x rune = 'a' //rune is an alias for int32
	var y int32 = 'b'

	println(x)
	println(y)
	println(string(x))
	println(string(y))
}
