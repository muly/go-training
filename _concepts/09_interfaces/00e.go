// an interface can be implemented by more than one type
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

func main() {
	var sh shape
	sh = square{ // square implements shape interface, because square has the area and perimeter methods implemented with the same signature.
		length: 10,
	}
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())

	sh = rectangle{ // rectangle also implements shape interface
		length: 10,
		width:  5,
		color:  "green",
	}

	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
}
