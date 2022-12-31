package main

import "fmt"

func main() {
	// 7forLoop/6range_withValueIgnoredInSimpleForm.go
	// range: ignoring the value; simplified form
	a := [4]int{9, 8, 7, 6}
	for i := range a {
		fmt.Println(a[i])
	}
}
