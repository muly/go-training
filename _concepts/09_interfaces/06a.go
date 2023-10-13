// type assertion

// type casting: lets look into type casting first

package main

import "fmt"
func main() {
	// START OMIT
	var i int = 10
	var f float32 = 12.34
	fmt.Println(i, f)

	// var j int = f // ERROR: cannot use f (variable of type float32) as int value in variable declaration
	// fmt.Println(i, j)

	var e float32
	// e = i // ERROR: cannot use i (variable of type int) as float32 value in assignment
	e=float32(i)
	fmt.Println(f, e)

	var s string = "100"
	// var k int = int(s) // ERROR: cannot convert s (variable of type string) to type int

// END OMIT
}
