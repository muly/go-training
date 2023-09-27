// an type can implement more than one interface

package main

import "fmt"
// START OMIT
type shape interface {
	area() float32
	perimeter() float32
}
type colorType interface {
	getColor() string
}
//////////////////////
type square struct {
	length float32
	color  string
}

func (s square) area() float32 {
	return s.length * s.length
}
func (s square) perimeter() float32 {
	return s.length * 4
}
func (s square) getColor() string {
	return s.color
}

// END OMIT

func main() {
	var s shape
	s = square{
		length: 10,
		color:  "red",
	}

	var c colorType
	c = square{
		length: 10,
		color:  "red",
	}

	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println(c.getColor())
}
