package main

import (
	"fmt"
)

type student struct {
	name    string
	marks   int
	address // embedding
	// address address // not embedding
	// schoolAddress address // not embedding
}

type address struct {
	stName string
	city   string
}

func main() {
	var s1 student = student{
		name:    "A",
		marks:   50,
		address: address{"1 main st", "AAA"},
	}
	fmt.Println(s1.address.city)
	fmt.Println(s1.city)
}
