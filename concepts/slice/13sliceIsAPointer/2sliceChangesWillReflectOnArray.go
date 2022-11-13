package main

import "fmt"

func main() {
	var a [4]int
	a = [4]int{1, 2, 3, 4}

	s := a[0:2]

	fmt.Println("before the change")
	fmt.Printf(" %v\n", a)
	fmt.Printf(" %v\n", s)

	s[0] = 100

	fmt.Println("after the change")
	fmt.Printf(" %v\n", a)
	fmt.Printf(" %v\n", s)
}
