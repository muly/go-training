// type assertion

package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	var j int
	// j = i // ERROR: cannot use i (variable of type interface{})
	// 				as int value in assignment: need type assertion
	j = i.(int) // type assertion

	// another example
	i = 12.34
	var k float64
	k = i.(float64)

	fmt.Println(j, k)
}
