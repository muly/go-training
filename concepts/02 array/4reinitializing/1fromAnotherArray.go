package main

import "fmt"

func main() {
	// 4reinitializing/1fromAnotherArray.go
	// reinitializing the array
	a := [4]int{1, 2, 3, 4}

	a = [4]int{10, 30, 50, 60}
	fmt.Println(a)
}
