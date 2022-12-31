package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	email     string
	english   int
	maths     int
}


// function
func fullname(s student) string {
	return s.firstName + " " + s.lastName
}

// method:
func (s student)fullname() string {
	return s.firstName + " " + s.lastName
}

func main() {
	s := student{firstName: "A",
		lastName: "L",
	}
	fmt.Println(s.fullname())
}
