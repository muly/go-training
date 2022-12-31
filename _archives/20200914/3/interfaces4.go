// type assertion

package main

import "fmt"

type math interface {
	isEven(y int) bool
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
func main() {
	var m math
	var v basicmath = basicmath{"my basic math"}
	m = v
	nbr := 10
	fmt.Println(nbr, m.isEven(nbr))

	process(v)
}

func process(m math) {
	var b basicmath
	b = m.(basicmath) // type assertion
	fmt.Println(b)

	var i int64
	var k int32
	k = int32(i) // type casting

	fmt.Println(i, k)

	// 1. convert it back to basicmath

}
