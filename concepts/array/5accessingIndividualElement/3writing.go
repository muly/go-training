package main

import "fmt"

func main() {
	//5accessingIndividualElement/3writing.go
	// changing individual element:
	a := [4]int{1, 2, 3, 4}
	a[1] = 100
	fmt.Println(a)
}
