package main

import "fmt"

func main() {
	var a [6]int
	a = [6]int{1, 2, 3, 4, 5, 6}

	s1 := a[2:]

	fmt.Printf("type = %T, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	fmt.Println(s1)
}
