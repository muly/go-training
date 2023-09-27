
package main

import "fmt"

// START OMIT
type shape interface {
	area() float32
	perimeter() float32
}

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
func (s square) getColor() string { // extra method which is not in interface definition 
	return s.color
}
// END OMIT

func main() {
	var sq square
	sq = square{ // square implements shape interface, because square has the area and perimeter methods implemented with the same signature.
		length: 10,
		color:  "blue",
	}
	fmt.Println(sq.area())
	fmt.Println(sq.perimeter())
	fmt.Println(sq.getColor()) // not an error

	var sh shape
	sh = square{ // square implements shape interface, because square has the area and perimeter methods implemented with the same signature.
		length: 10,
		color:  "blue",
	}
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
	// fmt.Println(sh.getColor()) // ERROR: sh.getColor undefined (type shape has no field or method getColor)
}
