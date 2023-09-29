// empty interface

package main

import "fmt"

// START OMIT
type rectangle struct {
	length float32
	width  float32
}
type square struct {
	length float32
}
type canon struct {
	name string
}

// END OMIT

func main() {
	var a interface{}

	a = square{length: 10}
	fmt.Println(a)

	a = rectangle{length: 10, width: 5}
	fmt.Println(a)

	a = canon{name: "CANON PRINTER"}
	fmt.Println(a)

	a = 10
	fmt.Println(a)

	a = map[string]string{"hello": "world"}
	fmt.Println(a)
}
