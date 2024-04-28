package main

import "fmt"

func main() {

	// START OMIT
	var m = map[int]string{1: "one", 3: "three", 0: "zero"}

	m = map[int]string{10: "ten", 20: "twenty", 30: "thirty"}
	fmt.Println(m)
	// END OMIT
	fmt.Println(m)
}
