// Nil interface value

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
	var s shape
	
	// s = square{ /// NOTE: s is not initialized
	// 	length: 10,
	// }

	fmt.Println(s) // nil
	fmt.Println(s.area()) // PANIC: nil pointer error
}
