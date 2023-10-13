// type assertion

package main

import "fmt"

func main() {
	// START OMIT
	var i interface{}
	i = 10

	var s string
	s = i.(string) 
	// END OMIT
	// PANIC: interface conversion: interface {} is int, not string

	fmt.Println(s)
}
