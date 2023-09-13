// function with interface type as a parameter

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
	color  string
}

func (s rectangle) area() float32 {
	return s.length * s.width
}
func (s rectangle) perimeter() float32 {
	return (s.length + s.width) * 2
}
func (s rectangle) getColor() string {
	return s.color
}

func describeSquare(s square) {
	fmt.Println("square data: ", s)
	fmt.Println("square area: ", s.area())
	fmt.Println("square perimeter: ", s.perimeter())
	fmt.Println("--------")
}

func describeShape(s shape) {
	fmt.Println("shape data: ", s)
	fmt.Println("shape area: ", s.area())
	fmt.Println("shape perimeter: ", s.perimeter())
	fmt.Println("--------")
}

func main() {
	sq := square{length: 10}
	r := rectangle{length: 10, width: 20}

	describeSquare(sq)
	// describeSquare(r) // ERROR: cannot use r (variable of type rectangle) as square value in argument to describeSquare

	describeShape(sq)
	describeShape(r)
}
