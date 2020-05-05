package main

import (
	"fmt"
)

func main() {
	var a int = 6
	double := func() {
		a=a*2
	}

	double()
	fmt.Println(a)

	double()
	fmt.Println(a)
}
