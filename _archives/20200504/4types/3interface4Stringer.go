package main

import (
	"fmt"
)

type myInt int

func (d myInt) String() string {
	return fmt.Sprintf("my value is %s\n", d)
}

type dog struct {
	name    string
	owner   string
	address string
}

func (d dog) String() string {
	return fmt.Sprintf("Dog Name: %s \nOwner Name: %s\n", d.name, d.owner)
}

func main() {
	var d dog
	d = dog{name: "max", owner: "owner1"}
	fmt.Println(d)

	var i myInt = 10
	fmt.Println(i)
}
