package main

import (
	"fmt"
)

type student struct {
	name  string
	email string
}

func main() {
	s := student{name: "student2", email: "dsdss"}
	s2 := s

	s2.name= "studenta"

	if s == s2 {
		fmt.Println("same")
	} else {
		fmt.Println("NOT same")
	}

}
