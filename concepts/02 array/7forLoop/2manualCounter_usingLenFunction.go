package main

import "fmt"

func main() {
	// 7forLoop/2manualCounter_usingLenFunction.go
	// for loop: using len()
	a := [4]int{1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
