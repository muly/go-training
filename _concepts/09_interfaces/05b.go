// empty interface

package main

import "fmt"

func main() {

	// type empty interface{} // No need to define a type for empty interface
	// var a empty
	var a interface{} // No need to define a type for empty interface

	a = square{length: 10}
	fmt.Println(a)

	a = canon{name: "CANON PRINTER"}
	fmt.Println(a)

	a = 10
	fmt.Println(a)
}
