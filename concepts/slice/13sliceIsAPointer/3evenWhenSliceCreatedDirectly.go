package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4}

	fmt.Printf(" %p\n", &s)

	fmt.Printf(" %p\n", &s[0]) // so, this should be the pointer to the underlying array
}
