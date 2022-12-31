package main

import "fmt"

func doubleGenerator(j int) func() int { // returns a function
	return func() int {
		j = j * 2
		//fmt.Println("inside child function", j)
		return j
	}
}

func main() {
	i := 10
	double1 := doubleGenerator(i)
	double2 := doubleGenerator(i)
	fmt.Println("outside child function", i)
	fmt.Println("outside child function", double1())  // 20
	fmt.Println("outside child function", double2())  // 20
	// fmt.Println("outside child function", double()) // ????????
	fmt.Println("outside child function", i)
}