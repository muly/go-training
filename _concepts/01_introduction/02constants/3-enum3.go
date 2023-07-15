package main

import "fmt"

// START OMIT
func main() {

	const (
		colorRed   = iota+1
		colorBlue  
		colorGreen 
	)
	fmt.Println(colorRed, colorBlue, colorGreen)
}

// END OMIT
