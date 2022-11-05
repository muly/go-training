package main

import "fmt"

func main() {
	// 7forLoop/3range_withIndexAndValue.go
	// range: intro
	a := [4]int{9, 8, 7, 6}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
