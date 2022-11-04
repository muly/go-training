//
package main

func main() {

	// //
	// { // initializing: declare and initialize
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// //
	// { // initializing: declare and initialize in same line
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// //
	// { // initializing: declare and initialize in same line, inferring the datatype
	// 	var a  = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// //
	// { // initializing: short hand form
	// a  := []int{1, 2, 3, 4}
	// fmt.Println(a)
	// }

	// //
	// { // length:
	// var a []int = []int{1, 2, 3, 4}
	// 	fmt.Println(len(a))
	// }

	// //
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

	// //
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

	// //
	// { // re-initializing the slice:
	// var a []int = []int{1, 2, 3, 4}

	// a =  []int{20, 30, 40, 50}
	// 	fmt.Println(a)
	// }

	// //
	// { // re-initializing the slice: with different length
	// var a []int = []int{1, 2, 3, 4}

	// a =  []int{20, 30, 40, 50, 60, 70}
	// 	fmt.Println(a, len(a))
	// }

	// //
	// { // un initialized slice: initialized with 0 length
	// var a []int
	// 	fmt.Println(a, len(a))
	// }

	// //
	// { // append: a single element
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	a = append(a, 5)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// //
	// { // append: multiple elements
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	a = append(a, 5, 6, 7)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// //
	// { // append: using another slice contents (... expansion notation)
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	var b = []int{5, 6, 7}

	// 	a = append(a, b...)
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// //
	// { // append: skipping assignment
	// 	var a []int = []int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  %v\n", len(a), a)

	// 	append(a, 5) // ERROR: append(a, 5) (value of type []int) is not used
	// 	fmt.Printf("len=%d  %v\n", len(a), a)
	// }

	// //
	// { // append: growing beyond the capacity of the slice: TODO: need to remonstrate after the concept of underlying array, and capacity topics are covered
	// }

	// //
	// { // capacity of a slice: intro
	// 	// capacity of a slice is the maximum number of elements the slice can hold.
	// 	// if we want to add more elements beyond the capacity, a new slice needs to be created.
	// 	// append() function will do this automatically for us.
	// 	//
	// 	// in the below example, the length and capacity is same because of the way we are initializing it.
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}
	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)
	// }

	// //
	// { // capacity of a slice: growing beyond: Note: while explaining this topic, I need to add references to address of the underlying array and how it changes when we grow a slice beyond its current capacity.

	// //
	// { // make(): with length and capacity
	// 	var a []int
	// 	a = make([]int, 2, 4)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// //
	// { // make(): with capacity only
	// 	var a []int
	// 	a = make([]int, 0, 4)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// //
	// { // make(): with length only
	// capacity will be defaulted to the length
	// 	var a []int
	// 	a = make([]int, 2)

	// 	fmt.Println(len(a), cap(a))
	// 	fmt.Println(a)
	// }

	// //
	// { // make and append: make with 0 length
	// 	var a []int
	// 	a = make([]int, 0, 4)

	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	// 	a = append(a, 50)
	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)
	// }

	// //
	// { // make and append: make with non zero length
	// 	var a []int
	// 	a = make([]int, 2, 4)

	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	// 	a = append(a, 50)
	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	// 	a[1] = 60 // you can always change the content of the initial cells
	// 	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)
	// }

	// //
	// { // slicing an array: - using both a low and high bounds
	// 	var a [6]int
	// 	a = [6]int{1, 2, 3, 4, 5, 6}

	// 	s1 := a[2:4]

	// 	fmt.Printf("type = %T, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	// 	fmt.Println(s1)
	// }

	// //
	// { // slicing an array: - using a low bound only
	// 	var a [6]int
	// 	a = [6]int{1, 2, 3, 4, 5, 6}

	// 	s1 := a[2:]

	// 	fmt.Printf("type = %T, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	// 	fmt.Println(s1)
	// }

	// //
	// { // slicing an array: - using a high bound only
	// 	var a [6]int
	// 	a = [6]int{1, 2, 3, 4, 5, 6}

	// 	s1 := a[:4]

	// 	fmt.Printf("type = %T, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	// 	fmt.Println(s1)
	// }

	// //
	// { // slicing an array: - using no bounds
	// 	var a [6]int
	// 	a = [6]int{1, 2, 3, 4, 5, 6}

	// 	s1 := a[:]

	// 	fmt.Printf("type = %T, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	// 	fmt.Println(s1)
	// }

	// //
	// { // slicing a slice:
	// 	var a [10]int
	// 	a = [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	// 	s1 := a[2:8]

	// 	s2 := s1[1:3]

	// 	fmt.Printf("type = %T, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	// 	fmt.Println(s2)
	// }

	// //
	// { // changing individual elements:
	// 	var s = []int{1, 2, 3, 4}
	// 	fmt.Println(s)

	// 	s[0] = 100
	// 	fmt.Println(s)
	// }

	// //
	// { // slice is a pointer to underlying array: slicing a array
	// 	// to demonstrate that lets slice an array and see the metadata for the created slice
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}
	// 	fmt.Printf(" %p\n", &a)

	// 	s := a[0:2]
	// 	fmt.Printf(" %p\n", &s)

	// 	fmt.Printf(" %p\n", &s[0]) // this gives the address of the first element of slice, which is the address of the underlying array
	// }

	// //
	// { // slice is a pointer to underlying array: changes to the slice are reflected in the underlying array
	// 	// to demonstrate that lets slice an array
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}

	// 	s := a[0:2]

	// 	fmt.Println("before the change")
	// 	fmt.Printf(" %v\n", a)
	// 	fmt.Printf(" %v\n", s)

	// 	s[0] = 100

	// 	fmt.Println("after the change")
	// 	fmt.Printf(" %v\n", a)
	// 	fmt.Printf(" %v\n", s)
	// }

	// //
	// { // slice is a pointer to underlying array: creating a slice directly
	// 	// lets create slice without explicitly creating an array
	// 	var s = []int{1, 2, 3, 4}

	// 	fmt.Printf(" %p\n", &s)

	// 	fmt.Printf(" %p\n", &s[0]) // so, this should be the pointer to the underlying array
	// }

	// //
	// { // copy() internal function: TODO:
	// }

	// //
	// { // for: counter variable, length hardcoded  TODO:
	// }

	// //
	// { // for: counter variable, using len() function  TODO:
	// }

	// //
	// { // for: using range  TODO:
	// }

	/////////////////////////////////////////////////////////////////////

	// { // delete an entry in slice: no direct function, we need to stitch the two slices, while ignoring the element that needs to be deleted.
	// 	var a = []int{1,2,3,4,5,6,7,8,9,0}
	// 	fmt.Println(a)
	// 	a= append(a[:4],a[5:]...)
	// 	fmt.Println(a, len(a), cap(a))
	// }

}
