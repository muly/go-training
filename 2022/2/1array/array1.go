//
package main

import "fmt"

func main() {

	// { // initializing: declare and initialize
	// 	var a [4]int
	// 	a = [4]int{1,2,3,4}
	// 	fmt.Println(a)
	// 	print(&a)

	// 	// fmt.Println(a[0:2])
	// 	// fmt.Println(a[4])
	// }

	// { // initializing: declare and initialize in same line
	// 	var a [4]int = [4]int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// { // initializing: declare and initialize in same line, inferring the datatype
	// 	var a = [4]int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }




	// { // initializing: short hand form
	// 	a := [4]int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// { // assigning one array to another array
	// 	a := [4]int{1, 2, 3, 4}
	// 	var b [4]int
	// 	b = a
	// 	fmt.Println(a, b)
	// }

	// { // assigning one array to another array of different length. array has the length as part of the data type
	// 	a := [4]int{1, 2, 3, 4}
	// 	var b [5]int
	// 	b = a
	// 	fmt.Println(a, b)
	// }

	// { // reinitializing the array 
	// 	a := [4]int{1, 2, 3, 4}

	// 	a = [4]int{10,30,50,60}
	// 	fmt.Println(a)
	// }

	// { // reinitializing the array: with different length
	// 	a := [4]int{1, 2, 3, 4}

	// 	a = [4]int{10,30,50,60}
	// 	fmt.Println(a)

	// 	a = [5]int{10,30,50,60, 70} // ERROR: 
	// 	fmt.Println(a)
	// }

	// // un initialized variable
	// var i int
	// fmt.Println(i)

	// { // un initialized array: initialized all elements with zero value
	// 	var a [4]int
	// 	fmt.Println(a) // [0 0 0 0]
	// }

	// { // len and cap system functions
	// 	a := [4]int{1, 2, 3, 4}
	// 	fmt.Println(len(a), cap(a))
	// }

	// { // accessing individual element
	// 	a := [4]int{1, 2, 3, 4}
	// 	fmt.Println(a[2])
	// 	fmt.Println(a[0])
	// 	fmt.Println(a[3])
	// }

	// { // accessing individual element: outside the boundary
	// 	a := [4]int{1, 2, 3, 4}
	// 	// fmt.Println(a[4]) //ERROR: invalid argument: array index 4 out of bounds [0:4]
	// }

	// { // changing  individual element:
	// 	a := [4]int{1, 2, 3, 4}
	// 	a[1]=100
	// 	fmt.Println(a)
	// }

	// { // for loop: length is hard coded
	// 	a := [4]int{1, 2, 3, 4}
	// 	for i:=0; i<4; i++ {
	// 		fmt.Println(a[i])
	// 	}
	// }

	// { // for loop: using len()
	// 	a := [4]int{1, 2, 3, 4}
	// 	for i:=0; i<len(a); i++ {
	// 		fmt.Println(a[i])
	// 	}
	// }

	// { // range: intro
	// 	a := [4]int{9, 8, 7, 6}
	// 	for i, v := range a {
	// 		fmt.Println(i, v)
	// 	}
	// }

	// { // range: ignoring the index
	// 	a := [4]int{9, 8, 7, 6}
	// 	for _, v := range a {
	// 		fmt.Println(v)
	// 	}
	// }

	// { // range: ignoring the value
	// 	a := [4]int{9, 8, 7, 6}
	// 	for i, _ := range a {
	// 		fmt.Println(a[i])
	// 	}
	// }

	// { // range: ignoring the value; simplified form
	// 	a := [4]int{9, 8, 7, 6}
	// 	for i := range a {
	// 		fmt.Println(a[i])
	// 	}
	// }

}
