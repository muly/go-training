// empty interface value

package main

import "fmt"

// 2 method interface
type shape interface {
	area() float32
	perimeter() float32
}

// implementing type must have those 2 matching methods
type square struct {
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}
func (s square) perimeter() float32 {
	return s.length * 4
}

// single method interface
type printer interface {
	print(string)
}

// implementing type must have that 1 matching method
type canon struct {
}

func (c canon) print(s string) {
	fmt.Println(s)
}

// 0 method interface
type any interface{}

// implementing type must have 0 matching methods, i.e, all types will implement the empty interface

func main() {
	var s shape
	s = square{
		length: 10,
	}
	fmt.Println(s)

	var p printer
	p = canon{}

	fmt.Println(p)

	var a any

	a = square{
		length: 10,
	}

	a = canon{}

	a = 10
	fmt.Println(a)

}

// 2 method interface
// implementing type must have those 2 matching methods

// single method interface
// implementing type must have those 1 matching method

// 0 method interface
// implementing type must have 0 matching method
