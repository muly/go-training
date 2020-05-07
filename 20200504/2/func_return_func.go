package main

import (
	"fmt"
)

type myAdd func(a, b int) (int, int)

func main() {
	f := generate()
	s, a := f(3, 10)
	fmt.Println(s, a)
}

func generate() myAdd {
	var f myAdd = func(a, b int) (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}

	return f
}
