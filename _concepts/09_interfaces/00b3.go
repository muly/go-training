package main

import "fmt"

// START OMIT
type shape interface {
	message() string
}

type myMap map[string]string // <--- this is user defined type

func (s myMap) message() string {
	return "hello from my map"
}

// END OMIT
func main() {
	var m myMap = map[string]string{}

	var v shape
	v = m

	fmt.Println(v.message())
}
