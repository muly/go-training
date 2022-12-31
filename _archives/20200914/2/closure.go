package main

import "fmt"

func main() {
	i := 10

	double := func() {
		i = i * 2
		fmt.Println("inside function ", i) // 
	}

	fmt.Println("outside function", i) //
	double()
	fmt.Println("outside function", i) //
	double()
	fmt.Println("outside function", i) //
	double()
	fmt.Println("outside function", i) // ???????????
}
