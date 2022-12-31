package main

import "fmt"

type math interface {
	isEven(y int) bool
	isOdd(y int) bool
}

type basicmath struct {
	name string
}



func (b basicmath) isEven(x int) bool {
	fmt.Println(b.name)
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func (b basicmath) isOdd(x int) bool {
	if x%2 != 0 {
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
	fmt.Println(nbr, m.isEven(nbr))


}
