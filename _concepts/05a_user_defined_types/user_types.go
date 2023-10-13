package main

import "fmt"

type myString string // user defined type
type studentMap map[string]int

func main() {

	var a string // string is golang defined type
	a = "hello"

	var b myString // variable from user defined type
	b = "hello"

	var s studentMap 
	s= map[string]int{"a":1}

	fmt.Println(a, b, s)

}
