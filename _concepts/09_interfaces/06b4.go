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
	_, ok := i.(int)
	if ok {
		processInt(i.(int))
	}

	f, ok := i.(float64)
	if ok {
		processFloat64(f)
	}
}

func processInt(x int) {
	// ..... does some stuff with int data
	fmt.Println("processing int", x)
}
func processFloat64(x float64) {
	// ..... does some stuff with float data
	fmt.Println("processing float64", x)
}