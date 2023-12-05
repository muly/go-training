package main

import "fmt"

// START OMIT
func main() {

	const ( // using iota enumerator
		colorRed   = 1
		colorBlue  //  = 1
		colorGreen //  = 1
	)
	fmt.Println(colorRed, colorBlue, colorGreen)
}

// END OMIT
