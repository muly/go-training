package main

import (
	"fmt"
)

type student struct {
	name          string
	email         string
	english       int
	maths         int
	homeaddress   address
	schooladdress address
}
type address struct {
	line1, line2 string
	city         string
	zipcode      int
}

func main() {
	s := student{
		name:        "student1",
		email:       "dsdss",
		homeaddress: address{city: "bangalore", zipcode: 123456}}

	s.english = 100
	s.homeaddress = address{city: "bangalore", zipcode: 123456}

	fmt.Println(s.name)
}
