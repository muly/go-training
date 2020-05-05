package main

import (
	"fmt"
)

type student struct {
	fname  string
	lname  string
	marks int
}

type myAdd func(a, b int) (int, int)

func main() {
	var s1 student = student{
		fname:  "F",
		lname:  "L",
		marks: 50,
	}
	fmt.Println(s1.fullName())

}

func (s student) fullName() string{
	return s.fname+" "+s.lname
}
