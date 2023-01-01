package main

import "fmt"

func main() {
	var a [4]int
	a = [4]int{1, 2, 3, 4}
	fmt.Printf(" %p\n", &a)

	s := a[0:2]
	fmt.Printf(" %p\n", &s)

	fmt.Printf(" %p\n", &s[0]) // this gives the address of the first element of slice, which is the address of the underlying array
}
