// type switch

package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	convert(i)

	i = 12.34
	convert(i)

	i="hello"
	convert(i)
}

func convert(i interface{}) {

	switch  i.(type){
	case int:
		processInt(i.(int))
	case float64:
		processFloat64(i.(float64))
	default:
		fmt.Printf("unsupported type %T\n", i)
	}

}

func processInt(x int) {
	fmt.Println("processing int", x)

}
func processFloat64(x float64) {
	fmt.Println("processing float64", x)
}
