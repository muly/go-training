// TODO: need to add the structs concepts with examples

package main

import "fmt"

type student struct {
	name    string
	address string
	marks   int
}

func main() {
var marks int
	fmt.Println("enter marks:")
	fmt.Scanf("%d", &marks)

	var s student = student{
		name:    "rohit",
		address: "hyderabad",
	}

	s.marks = marks

	fmt.Println(s)
}
