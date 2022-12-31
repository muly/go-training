package main

import "fmt"

func main() {
	// 4reinitializing/2fromAnotherArrayOfDifferentLength.go
	// reinitializing the array: with different length
	a := [4]int{1, 2, 3, 4}

	a = [4]int{10, 30, 50, 60}
	fmt.Println(a)

	a = [5]int{10, 30, 50, 60, 70} // ERROR:
	fmt.Println(a)
}
