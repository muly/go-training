package main

import (
	"fmt"
)

type pet interface {
	showDetail()
	// showDetail(string) // Note: overloading not allowed. ERROR: ambiguous selector d.showDetail
}

type renamablePet interface {
	showDetail()
}

type dog struct {
	name    string
	owner   string
	address string
}

func (d dog) showDetail() {
	fmt.Println(d)
}

func (d dog) showDetail(s string) {
	fmt.Println(s, d)
}

func main() {
	var d pet
	d = &dog{name: "max"}
	d.showDetail()

	fmt.Println(d)

	var r renamablePet
	r = &dog{name: "max"}
	r.showDetail()

}
