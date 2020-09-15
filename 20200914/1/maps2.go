package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	m1["student1"]=120
	m2=m1

	m2["student1"]=90
	fmt.Println(m1, m2)
}
