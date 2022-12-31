package main

import (
	"fmt"
)

func main() {
	age := 30

	switch {
	case age < 15:
		fmt.Println("cannot drive")
	case age < 30:
		fmt.Println("can drive")
	default:
		fmt.Println("can also drive")
	}
}
