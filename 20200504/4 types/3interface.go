package main

import (
	"fmt"
)

type pet interface {
	changeOwner(string)
	// changeAddress(string) // ERROR: *dog does not implement pet (missing changeAddress method)
	showDetail()
}

type dog struct {
	name  string
	owner string
	address string
}

func (d *dog) changeOwner(newOwner string) {
	d.owner = newOwner
}

func (d *dog) changeName(newName string) { //NOTE: dog can have other methods not in interface, but still implement the interface
	d.name = newName
}

// func (d dog) showdetail() { // ERROR: name of the function along with the input/output signature should be same as in the interface, for this struct to implement it
func (d dog) showDetail() {
	fmt.Println(d)
}



func main() {
	var d pet
	d = &dog{name: "max"}

	fmt.Println(d)
}
