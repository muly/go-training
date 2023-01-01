package main

import "fmt"

func main() {
	// 7forLoop/4range_withIndexIgnored.go
	// range: ignoring the index
	a := [4]int{9, 8, 7, 6}
	for _, v := range a {
		fmt.Println(v)
	}
}
