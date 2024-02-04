package main

import "fmt"

func main() {
	var m map[int]string
	m = map[int]string{1: "one", 3: "three", 0: "zero"}

	fmt.Println(m)
}
