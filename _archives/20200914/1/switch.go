package main

import (
	"fmt"
)

var i = 50

func init(){

}

func main() {
	i := 50
	switch {
	case (i<=30):
		fmt.Println("value is 30")
	case i==50:
		fmt.Println("value is 50")
	case i==60:
		fmt.Println("value is 60")
	case i==70:
		fmt.Println("value is 70")
	default:
		fmt.Println("default case; value is 50")
	}

}
