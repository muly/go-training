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

	a = square{length: 10} // assigning square type to empty interface
	fmt.Println(a)

	a = rectangle{length: 10, width: 5} // rectangle type
	fmt.Println(a)

	a = canon{name: "CANON PRINTER"} // canon type
	fmt.Println(a)

	a = 10 // INT type
	fmt.Println(a)

	a = map[string]string{"hello": "world"} // MAP
	fmt.Println(a)
}
