package main

import "fmt"

type doubleType func()

func generator(d doubleType){
	d()
}

func main() {
	i := 10

	double := func() {
		i = i * 2
		fmt.Println("inside function ", i) // 
	}

	generator(double)
	fmt.Println("outside function", i) //
	generator(double)
	fmt.Println("outside function", i) //
}
