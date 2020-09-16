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
func fullname(s *student) string {
	return s.firstName + " " + s.lastName
}

// method:
// func (s student)update(f, l string) {
// 	 s.firstName = f
// 	 s.lastName = l
// 	 fmt.Println(s)
// }

func (s student)fullname() string {
	return s.firstName + " " + s.lastName
}

// type myInt int
// func (i *myInt) double(){
// 	*i=*i+*i
// }

func (s *student)update(f, l string) {
	s.firstName = f
	s.lastName = l
	fmt.Println(s.fullname())
}

func main() {
	s := student{firstName: "A",
		lastName: "L",
	}

	s.update("a", "l")
	fmt.Println(s)
}
