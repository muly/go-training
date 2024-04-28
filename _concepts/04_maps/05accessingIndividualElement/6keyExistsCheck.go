package main

import "fmt"

func main() {

	// START OMIT
	// {
	var m = map[int]float32{1: 10.5, 3: 30.5, 5: 0.0}

	fmt.Println(m[99], m[5])
	v, ok := m[99]
	fmt.Printf("value: %v, key exists: %v\n", v, ok)

	fmt.Println()
	v, ok = m[5]
	fmt.Printf("value: %v, key exists: %v\n", v, ok)
	// }
	// END OMIT
	fmt.Println(m)
}
