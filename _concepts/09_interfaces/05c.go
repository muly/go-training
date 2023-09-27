// function with parameter as empty interface

package main

import (
	"fmt"
)

// START OMIT

func myPrint(a interface{}) {
	// do some stuff
	fmt.Println(a)
}

func main() {
	var s string = "golang" // string
	var i int = 999         // int
	var st = struct {       // struct
		s string
		i int
	}{s: "hello", i: 10}

	myPrint(s)
	myPrint(i)
	myPrint(st)
}

// END OMIT
