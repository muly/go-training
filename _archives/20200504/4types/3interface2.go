package main

import (
	"fmt"
)

type pet interface {
	changeOwner(string)
	showDetail()
}

// this interface has some methods same as the previous interface
type renamablePet interface {
	changeName(newName string)
	showDetail()
}

// this interface has same methods as pet interface
type random interface {
	changeOwner(string)
	showDetail()
}

// dog type implements both the pet and renamablePet interfaces (as well as the random interface)
type dog struct {
	name    string
	owner   string
	address string
}

func (d *dog) changeOwner(newOwner string) {
	d.owner = newOwner
}
func (d *dog) changeName(newName string) {
	d.name = newName
}
func (d dog) showDetail() {
	fmt.Println(d)
}

// but cat implements only the pet interface (and the random interface)
type cat struct {
	name    string
	owner   string
	address string
}

func (d *cat) changeOwner(newOwner string) {
	d.owner = newOwner
}
func (d cat) showDetail() {
	fmt.Println(d)
}

func main() {
	var d pet
	d = &dog{name: "max"}
	fmt.Printf("%T \n", d)
	d = &cat{name: "cat1"}

	fmt.Println(d)

	var r renamablePet
	r = &dog{name: "max"}
	fmt.Printf("%T \n", r)
	// r = &cat{name: "cat1"} // Note: cat doesn't implement renamablePet. ERROR: *cat does not implement renamablePet (missing changeName method)

	{
		var d random
		d = &dog{name: "max"}
		fmt.Printf("%T \n", d)
	}

}
