package main

import (
	"fmt"
)

type student struct {
	fname          string
	lname          string
	email         string
	english       int
	maths         int
	address
}
type address struct {
	line1, line2 string
	city         string
	zipcode      int
}

func (a address)fulladdress()string{
	return fmt.Sprintf("%s, %s, %s, %d", a.line1,a.line2,a.city,a.zipcode)
}

func main() {
	s := student{
		fname:        "student1",
		email:       "dsdss",
		address: address{city: "bangalore", zipcode: 123456}}

	s.english = 100
	s.address = address{city: "bangalore", zipcode: 123456}

	// fmt.Println(s.fullname())
	fmt.Println(s.fulladdress())
}
