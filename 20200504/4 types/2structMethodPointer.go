package main

import (
	"fmt"
)

type student struct {
	fname  string
	lname  string
	marks int
}

func main() {
	var s1 student = student{
		fname:  "F",
		lname:  "L",
		marks: 50,
	}

	fmt.Println(s1)
	s1.change("f1", "l1")
	fmt.Println(s1)
	s1.changePointer("f1", "l1")
	fmt.Println(s1)

}

func (s student) change(fname, lname string) {
	s.fname = fname
	s.lname = lname
}

func (s *student) changePointer(fname, lname string) {
	s.fname = fname
	s.lname = lname
}
