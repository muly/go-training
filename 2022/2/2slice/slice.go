//
package main

import "fmt"

func main() {

	//
	// { // initializing: declare and initialize
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// {// initializing: declare and initialize in same line
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// {// initializing: declare and initialize in same line, inferring the datatype
	// 	var a  = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// { // initializing: short hand form
	// a  := []int{1, 2, 3, 4}
	// fmt.Println(a)
	// }

	// { // length:
	// var a []int = []int{1, 2, 3, 4}
	// 	fmt.Println(len(a))
	// }

	//
	// { // assigning one slice to another slice: of same length
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}

	// 	var b []int
	// 	b = []int{20,30,40,50}
	// 	fmt.Println(len(a), len(b))

	// 	b = a

	// 	fmt.Println(len(a), len(b))
	// 	fmt.Println(a, b)
	// }

	// { // assigning one slice to another slice: of different length.
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}

	// 	var b []int
	// 	b = []int{20,30,40,50,60}
	// 	fmt.Println(len(a), len(b))

	// 	b = a

	// 	fmt.Println(len(a), len(b))
	// 	fmt.Println(a, b)
	// }

	//
	// { // re-initializing the slice:
	// var a []int = []int{1, 2, 3, 4}

	// a =  []int{20, 30, 40, 50}
	// 	fmt.Println(a)
	// }

	// { // re-initializing the slice: with different length
	// var a []int = []int{1, 2, 3, 4}

	// a =  []int{20, 30, 40, 50, 60, 70}
	// 	fmt.Println(a, len(a))
	// }

	//
	// { // un initialized slice: initialized with 0 length
	// var a []int
	// 	fmt.Println(a, len(a))
	// }

	//
	// { // append: a single element
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	a = append(a, 5)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// { // append: multiple elements
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	a = append(a, 5, 6, 7)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// { // append: using another slice contents (... expansion notation)
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	var b = []int{5, 6, 7}

	// 	a = append(a, b...)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// { // append: skipping assignment
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	append(a, 5) // ERROR: append(a, 5) (value of type []int) is not used
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	//
	// { // slice and underlying array: 
	// TODO: need a good example to demonstrate this by printing the addresses of slice and an array. how to see the corresponding array details of a slice variable.
	// TODO: TODO:  TODO: example to show the slice points to an underlying array
	//
	// }

	//
	// { // capacity of a slice: intro
		
	// }

	// { // capacity of a slice: growing beyond: Note: while explaining this topic, I need to add references to address of the underlying array and how it changes when we grow a slice beyond its current capacity.

	// slicing an array: TODO: need to add sub topics

	// slicing a slice: TODO: need to add sub topics

	/////////////////////////////////////////////////////////////////////

	// {
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}

	// 	s1 := a[0:2]

	// 	fmt.Printf("%T\n",s1)
	// 	fmt.Println(s1, len(s1))

	// 	s2 := a[:2] // [0:2]
	// 	fmt.Println(s2, len(s2))// [1 2] 2

	// 	s3 := a[2:]
	// 	fmt.Println(s3, len(s3)) // [3 4] 2

	// 	s4 := a[:]
	// 	fmt.Println(s4, len(s4)) // [1 2 3 4] 4
	// }

	// { // append: todo: need better example to demonstrate change or memory location for a slice when it grows beyond capacity.
	// 	var s string = "some value"
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}
	// 	print(&a)
	// 	print(&s)
	// 	fmt.Println(a, len(a))

	// 	a = append(a, 100, 200)
	// 	print(&a)
	// 	fmt.Println(a, len(a))
	// }

	// capacity:

	// {
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}
	// 	fmt.Println(len(a), cap(a))

	// 	s := a[0:2]
	// 	s2 := a[2:]
	// 	fmt.Println(s2)
	// 	fmt.Println(len(s), cap(s))

	// 	// print(&a)
	// 	// fmt.Println()
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)

	// 	s = append(s, 10, 20)
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)

	// 	s = append(s, 30)
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)
	// }

	// {
	// 	var a []int
	// 	var b []int
	// 	a = []int{1, 2, 3, 4}
	// 	b = a
	// 	fmt.Println(a)
	// 	fmt.Println(len(a), cap(a))
	// 	print(&a)

	// 	fmt.Println(b)
	// 	fmt.Println(len(b), cap(b))
	// 	print(&b)
	// 	fmt.Println()
	// }

	// {
	// 	a := []int{1, 2, 3, 4}
	// 	for _, b:= range a{
	// 		fmt.Println(b)
	// 	}
	// 	fmt.Println(len(a),cap(a))
	// }

	// {
	// 	var a []int
	// 	a = make([]int, 2, 4)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// {
	// 	var a []int
	// 	a = make([]int, 0, 4)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// {
	// 	var a []int
	// 	a = make([]int, 2)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// {
	// 	var a []int
	// 	a = make([]int, 2, 4)

	// 	fmt.Println(len(a),cap(a))
	// 	fmt.Println(a)

	// 	a = append(a, 50)
	// 	fmt.Println(a)

	// 	a[1]= 60
	// 	fmt.Println(a)
	// }

	// {
	// 	var a = []int{1,2,3,4,5,6,7,8,9,0}
	// 	fmt.Println(a)
	// 	a= append(a[:4],a[5:]...)
	// 	fmt.Println(a, len(a), cap(a))
	// }

	// {
	// 	var a = []int{}
	// 	for i:=0; i<100; i++{
	// 		a=append(a,i)
	// 		fmt.Println(len(a), cap(a))
	// 	}
	// }

	// {
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}

	// 	s := a[0:2]
	// 	s[0] = 99
	// 	fmt.Println(a)

	// 	for i := 0; i < 5; i++ {
	// 		s = append(s, i)
	// 	}
	// 	fmt.Println(s)
	// 	fmt.Println(a)
	// }

}
