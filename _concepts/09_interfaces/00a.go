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
func (s square) perimeter() float32 {
	return s.length * 4
}

// END OMIT

func main() {
	// using concrete type directly
	var sq square
	sq = square{
		length: 10,
	}
	fmt.Println(sq.area())
	fmt.Println(sq.perimeter())

	// using interface
	var sh shape
	sh = square{ // square implements shape interface, because square has the area and perimeter methods implemented (with the same signature as in the interface definition).
		length: 10,
	}
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
}