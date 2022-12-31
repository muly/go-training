package main

import (
	"fmt"
)

func main() {
	var a int
	p := &a // ref

	v := *p // de ref

	print(a)
	fmt.Println()
	print(p)
	fmt.Println()
	print(v)
	fmt.Println()

	{
		var a int
		var p *int
		p = &a
		print(a)
		fmt.Println()
		print(p)
		fmt.Println()
	}

}
