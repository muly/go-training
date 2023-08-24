package main

import (
	"fmt"
)

type customer struct {
	fname, lname, email string
}

// this is a function
func fullname(s customer) string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

// this is a method
func (s customer) fullname() string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

func main() {
	c := customer{fname: "Anitha", lname: "Muly", email: "E"}

	fmt.Println(fullname(c))  // calling a function
	fmt.Println(c.fullname()) // calling a method
}
