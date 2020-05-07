package main

import (
	"fmt"

	var2 "github.com/muly/go-training/20200504/3/vari"
)

func main() {
	s, _ := var2.Add(1, 2)
	fmt.Println(var2.Marks, s)

	fmt.Println(var2.AddAll("", 1, 2, 3, 4, 5))

	// exported function as type
	var f2 var2.MyAdd = func(a, b int) (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}

	// exported function var as type
	{
		su, av := var2.F(1, 6)
		fmt.Println(su, av)
	}
	{
		su, av := f2(10, 60)
		fmt.Println(su, av)
	}

	// exported struct
	{
		s := var2.Student{Fname: "f"}
		fmt.Println(s)

	}

	// implement the exported interface
	var d var2.Pet
	d = &dog{name: "max"}
	fmt.Printf("%T \n", d)

}

// // NOT ALLOWED
// func (s var2.Student) fullName() string{
// 	return s.fname+" "+s.lname
// }
