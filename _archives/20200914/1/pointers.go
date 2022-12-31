package main

import "fmt"

func main() {
	var v int = 10
	fmt.Println(v) // 10
	//&v address

	var p *int
	p = &v

	fmt.Println(*p) // 10
	*p = 12

	fmt.Println(v) // 12
}
