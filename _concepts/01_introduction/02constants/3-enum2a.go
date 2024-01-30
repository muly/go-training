package main

import "fmt"

// START OMIT
func main() {

	const ( // using iota enumerator
		colorRed   = iota
		colorBlue  = iota
		colorGreen = iota
	)
	fmt.Println(colorRed, colorBlue, colorGreen)
}

// END OMIT
