package main

import (
	"fmt"
)

func main() {
	
	// array
	// slice
	// map
	// struct

// var i int
// var j int8
// i := 10
// j = i

// var a [2]string
// a := [2]string{"student a", "student b"}
// var i int = 10
// var a [2]string = [...]string{"student a", "student b"}

// var a = [2][3]string{{"a", "b", "c"},{"c","d","e"}}

// // fmt.Println(a)
// a[0][0]="student a1"
// a[0][1]="student a2"
// a[2]=""
// fmt.Println(a[3])
a := [...]int{1,2,3}
fmt.Println(a)
fmt.Println(cap(a))
}