package main

import "fmt"

func main() {
	var m = map[int]int{1: 100, 2: 200}
	fmt.Println(m)
	insert(m, 3, 300) // note: we are passing the map variable without * (like we do in case of the pointer type),
	fmt.Println(m)    // still, the change made in the insert() function scope is visible outside the insert() function, i.e in main() function.
	// this behavior usually happens if we pass the variable as pointer. but since map is a pointer type, we have the same behavior without using *

}

func insert(m map[int]int, k, v int) {
	m[k] = v
}
