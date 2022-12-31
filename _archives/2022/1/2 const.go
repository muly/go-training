package main

import "fmt"


func main() {

	// not allowed:
	// const newPi float32
	// newPi = 3.414 //
	const pi float32 = 3.414
	const anotherPi = 3.414
	yetAnoterPi := 3.414 // this becomes a variable

	fmt.Print(class, pi, anotherPi, yetAnoterPi)
}
