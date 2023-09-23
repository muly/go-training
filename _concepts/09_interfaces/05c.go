// function with parameter as empty interface

package main

import (
	"fmt"
)



func myPrint(a interface{}) {
	// To some stuff
	fmt.Println(a)
}

func main() {
	var s string = "golang"
	var i int = 999
	var st = struct {
		s string
		i int
	}{s: "hello", i: 10}

	myPrint(s)
	myPrint(i)
	myPrint(st)
}
