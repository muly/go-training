package main

import "fmt"

type myString string // user defined type

func (s myString) prefix(p string) string {
	return fmt.Sprintf("%s%s", p, s)
}

// method on golang system type is not allowed
func (s string) prefix(p string) string { // ERROR: cannot define new methods on non-local type string
	return fmt.Sprintf("%s%s", p, s)
}

func main() {
	var b myString // variable from user defined type
	b = "world"

	fmt.Println(b.prefix("hello"))
}
