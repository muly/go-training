package main

import (
	"fmt"
)

// immutable string

func main() {
	str := "hello class"

	fmt.Println(str[0])
	// str[0]='x'

	str = "gello class"
	fmt.Println(string(str[0]))

}
