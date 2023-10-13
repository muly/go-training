package main

import "fmt"

// START OMIT
type shape interface {
	message() string
}

type myString string // <--- this is user defined type

func (s myString) message() string {
	return "hello from my string"
}

// END OMIT
func main() {
	var s myString

	var v shape
	v = s
	fmt.Println(v.message())

}
