package main

import (
	"fmt"
)



type student struct {
	name  string
	email string
	marks map[string]int
}

func main() {
s1 := student{name:"s1"}
s2 := student{name:"s2"}
	m1 := make(map[string]int)
	m1["english"]=120
	m1["math"]=90

	s1.marks= m1
	s2.marks = m1

s1.marks["english"]=110

fmt.Println(s1)
fmt.Println(s2)


}
