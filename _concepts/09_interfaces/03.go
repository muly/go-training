// interface value %v, and its concrete type %T

package main

import "fmt"

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

type rectangle struct {
	length float32
	width  float32
}

func (s rectangle) area() float32 {
	return s.length * s.width
}
func (s rectangle) perimeter() float32 {
	return (s.length + s.width) * 2
}

func main() {
	var s shape
	s = square{
		length: 10,
	}
	fmt.Printf("%T: %+v\n", s, s)

	s = rectangle{
		length: 10,
		width:  20,
	}
	fmt.Printf("%T: %+v\n", s, s)
}
