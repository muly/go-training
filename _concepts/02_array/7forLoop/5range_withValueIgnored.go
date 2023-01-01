package main

import "fmt"

func main() {
	// 7forLoop/5range_withValueIgnored.go
	// range: ignoring the value
	a := [4]int{9, 8, 7, 6}
	for i, _ := range a {
		fmt.Println(a[i])
	}
}
