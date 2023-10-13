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
	name string
}

func (c canon) print(s string) {
	fmt.Println(s)
}

// 0 method interface
type empty interface {
	// Note: no methods
}

// implementing type must have 0 matching methods, i.e, all types will implement the empty interface

func main() {
	// 2 method interface
	var s shape
	s = square{length: 10}
	fmt.Println(s)

	// 1 method interface
	var p printer
	p = canon{name: "CANON PRINTER"}

	fmt.Println(p)

	// empty interface
	var a empty
	
	a = square{length: 10}
	fmt.Println(a)

	a = canon{name: "CANON PRINTER"}
	fmt.Println(a)

	a = 10
	fmt.Println(a)
}
