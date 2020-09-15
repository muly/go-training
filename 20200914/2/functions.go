package main

import (
	"fmt"
)

var i = 50

// func init(){

// }

func main() {
	// var str  string
	// str := derive(i)
	// fmt.Println(str)

	fmt.Println(derive())
}

func derive()string{
	switch {
	case (i<=30):
		return "value is 30"
	case i==50:
		return "value is 50"
	case i==60:
		return "value is 60"
	case i==70:
		return "value is 70"
	default:
		return "default case; value is 50"
	}
}
