package main

import "fmt"

// START OMIT
func main() {

	const ( // using iota enumerator
		colorRed   = iota
		colorBlue  
		colorGreen 
	)
	fmt.Println(colorRed, colorBlue, colorGreen)
}

// END OMIT
