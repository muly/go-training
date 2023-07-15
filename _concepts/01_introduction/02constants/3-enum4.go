package main

import "fmt"


// START OMIT
func main() {

	const (
		colorRed   = iota+1
		colorBlue  
		colorGreen 
		
		colorOrange  = iota+4
		colorPink
	)
	fmt.Println(colorRed, colorBlue, colorGreen, colorOrange,colorPink )
}

// END OMIT
