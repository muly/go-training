package main

import "fmt"

// START OMIT
type shape interface {
	area() float32
	perimeter() float32
}

type square struct {
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}

func main() {
	var sh shape
	sh = square{ // ERROR: square does not implement shape (missing method perimeter)
		length: 10,
	}
	fmt.Println(sh.area())
}

// END OMIT
