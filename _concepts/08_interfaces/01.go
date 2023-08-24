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

type sphere struct {
	radius float32
}

func (s sphere) area() float32 {
	return 12345
}
func (s sphere) volume() float32 {
	return 34567
}

type line struct {
	length float32
}

func (l line) len() float32 {
	return l.length
}

func main() {
	var sh shape

	sh = square{ // square implements shape interface, because square has the area and perimeter methods implemented with the same signature.
		length: 10,
	}
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())

	sh = rectangle{
		length: 10,
		width:  5,
		color:  "green",
	}

	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())

	// sh = line{length: 10} // ERROR: line does not implement shape (missing method area)

	// sh = sphere{radius: 100} // ERROR: sphere does not implement shape (missing method perimeter)

}
