package main

import (
	"fmt"
)

type pet interface {
	showDetail()
}

type myS struct {
	pet
}

type dog struct {
	name    string
	owner   string
	address string
}

func (d dog) showDetail() {
	fmt.Println(d)
}

func main() {
	var d myS
	d.pet = &dog{name: "max"}
	d.showDetail()

	fmt.Println(d)
}
