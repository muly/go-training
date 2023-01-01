package main

import "fmt"

func main() {
	var a [10]int
	a = [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	s1 := a[2:8]

	s2 := s1[1:3]

	fmt.Printf("type = %T, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	fmt.Println(s2)
}
