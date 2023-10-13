// an type can implement more than one interface

package main

import "fmt"

// START OMIT
type shape interface {
	area() float32
	perimeter() float32
}

// ////////////////////
type square struct {
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}
func (s square) perimeter() float32 {
	return s.length * 4
}

// ////////////////////
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

// END OMIT

// func describeSquare(s square) {
// 	fmt.Printf("I'm a square, with data %+v\n", s)
// }

// func describeRectangle(s rectangle) {
// 	fmt.Printf("I'm a rectangle, with data %+v\n", s)
// }

func describeShape(s shape) {
	fmt.Printf("I'm a shape of type %T, with data %+v\n", s, s)
}

func main() {
	var s shape

	s = square{length: 10}
	describeShape(s)

	s = rectangle{length: 10, width: 5}
	describeShape(s)
}
