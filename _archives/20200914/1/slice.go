package main

import("fmt")

func main(){
	a := [4]int{2,4,5,6}
	v := 0
	fmt.Println(v)

	var s []int
	s= a[2:]
	
	fmt.Println(a)
// 	s[0]= 200
// fmt.Println(len(s), cap(s))

// s = a[0:4]
// fmt.Println(len(s), cap(s))
// fmt.Println(s[3])

fmt.Println(len(s), cap(s))
s = append(s, 500, 600, 700)
fmt.Println(len(s), cap(s))



	// fmt.Println(s, len(s), cap(s))

	// s2 := []int{100, 200, 300}

	// fmt.Println(s2, len(s2), cap(s2))
}