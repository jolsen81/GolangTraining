package main

import (
	"fmt"
	"math"
)

// type Triangle struct {
// 	side_1 float64
// 	side_2 float64
// 	side_3 float64
// }

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	diameter float64
}

func (c Circle) area() float64 {
	return (c.diameter / 2) * (c.diameter / 2) * math.Pi
}

type Shape interface {
	area() float64
}

func info(z ...Shape) {
	for _, s := range z {
		fmt.Printf("%T\n", s)
		fmt.Println(s)
		fmt.Println(s.area())
	}

}

func main() {

	s := Square{10}
	// t := Triangle{10, 5, 3}
	c := Circle{5}
	shapes := []Shape{s, c}
	// shapes[0] = s
	// shapes[1] = c
	info(shapes...)
}
