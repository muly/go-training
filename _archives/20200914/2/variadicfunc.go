package main

import (
	"fmt"
)

func main() {
	// a, b := 1, 2
	// fmt.Println(a, b)
	// fmt.Println(add("test",a, b)) 
	// fmt.Println(add("test 2 ",a)) 
	// fmt.Println(add("test 3 ")) 
	// fmt.Println(add("test",[]int{a, b})) 
	// fmt.Println(add("test 2 ",[]int{a})) 
	// fmt.Println(add("test 3 ", []int{})) 

	input := []int{1,2,3,4,5}
	fmt.Println(add("test",input...))
	fmt.Println(add("test",1,2,3,4,5))
	// Would fmt.Println(input..) print {1, 2, 3, 4} or 1,2,3,4
	fmt.Println()
}

// func add(name string, a []int)(int){
func add(name string, a ...int)(int){
	fmt.Printf("executing %s addition\n", name)
	sum := 0
	for _, v:= range a{
		sum+=v
	}
	return sum
}


