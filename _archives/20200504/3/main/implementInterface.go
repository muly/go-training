package main

import (
	"fmt"
)

type dog struct {
	name    string
	owner   string
	address string
}

func (d *dog) ChangeOwner(newOwner string) {
	d.owner = newOwner
}

func (d *dog) ChangeName(newName string) {
	d.name = newName
}

func (d dog) ShowDetail() {
	fmt.Println(d)
}
