// type switch

package main

import "fmt"

func main() {
	convert(10)

	convert(12.34)

	convert("hello")
}

func convert(i interface{}) {
	switch i.(type) {
	case int:
		processInt(i.(int))
	case float64:
		processFloat64(i.(float64))
	default:
		fmt.Printf("unsupported type %T\n", i)
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
