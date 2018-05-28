package main

import (
	"fmt"
	"log"
)

type SquareError struct {
	Val float64
	Err error
}

func (e *SquareError) Error() string {
	return fmt.Sprintf("norgate math error: %v %v", e.Val, e.Err)
}
func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math redux: square root of negative integer: %v", f)
		return 0, &SquareError{f, ErrNorgateMath}
	}
	return 42, nil
}
