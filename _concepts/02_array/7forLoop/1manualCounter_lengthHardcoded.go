package main

import "fmt"

func main() {
	// 7forLoop/1manualCounter_lengthHardcoded.go
	// for loop: length is hard coded
	a := [4]int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		fmt.Println(a[i])
	}
}
