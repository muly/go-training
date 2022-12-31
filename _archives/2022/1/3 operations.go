package main

import "fmt"

func main() {

	// arithematic operations
	var a, b int = 10, 2
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	// type casting
	var c float32 = 10.5
	fmt.Println(c+float32(a)) 
	// var s string = "10"
	// fmt.Println(c+float32(s))  // not allowed

	var s1, s2 string = "hello", "world"
	fmt.Println(s1+s2)

	var s3, s4 string = "10", "20"
	fmt.Println(s3+s4) // 1020


	fmt.Printf("%s %s\n", s1, s2)

	fmt.Printf("my numbers are %d,%d\n", a, b)
}

