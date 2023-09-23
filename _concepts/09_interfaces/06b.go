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
	// inside this function we want to do different things based on the concrete type
	fmt.Printf("%T\n", i)

	// processInt(i) // ERROR: cannot use i (variable of type interface{}) as int value in argument to processInt: need type assertion

	// processInt(i.(int))         // ERROR: panic: interface conversion: interface {} is float64, not int
	// processFloat64(i.(float64)) // ERROR: panic: interface conversion: interface {} is int, not float64

	// i.(int)

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
	fmt.Println("processing int", x)

}
func processFloat64(x float64) {
	fmt.Println("processing float64", x)
}
