package main

import (
	"fmt"

	var2 "training/2/vari"
)

func main() {
	s, _ := var2.Add(1, 2)
	fmt.Println(var2.Marks, s)

	fmt.Println(var2.AddAll("", 1, 2, 3, 4, 5))
}
