// empty interface

package main

import "fmt"

type math interface {
	// isEven(y int) bool
}

type basicmath struct {
	name string
}

func (b basicmath) isEven2(x int) bool {
	fmt.Println(b.name)
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var m math
	var v basicmath = basicmath{"my basic math"}
	m = v
	nbr := 10
	//fmt.Println(nbr, m.name, m.isEven2(nbr)) // compile time error
	fmt.Println(nbr, m.(basicmath).name, m.(basicmath).isEven2(nbr)) // now allowed using type assersion
}

