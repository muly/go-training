// type assertion

package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	convert(i)

	i = 12.34
	convert(i)
}

func convert(i interface{}) {
	fmt.Printf("%T\n", i)
	processInt(i.(int))
	processFloat64(i.(float64))
}

// START OMIT
func processInt(x int) {
	// ..... does some stuff with int data
	fmt.Println("processing int", x)
}
func processFloat64(x float64) {
	// ..... does some stuff with float data
	fmt.Println("processing float64", x)
}

// END OMIT
