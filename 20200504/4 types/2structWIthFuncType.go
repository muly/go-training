package main

import (
	"fmt"
)

type student struct {
	fname  string
	lname  string
	marks int
	fullName     full
}

type full func(a, b string) (string)

func main() {
	var f1 full = func(f, l string) string{
		return f+" "+l
	}	

	var s1 student = student{
		fname:  "F",
		lname:  "L",
		marks: 50,
		fullName:     f1,
	}
	fmt.Println(s1)

	fmt.Println(s1.fullName(s1.fname, s1.lname))
}
