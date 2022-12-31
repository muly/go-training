package main

import (
	"fmt"
	"strconv"
)

var isMale bool = true

var (
	marks int = 50
	name      = "go class"
)

const pi = 3.414

func main() {
	marks = 100

	isPass := true

	fmt.Println(isPass, isMale)

	address := "india" // shorthand
	// address = 200 //
	{
		var location = "123"
		var loc int
		loc, _ = strconv.Atoi(location)

		fmt.Println(location, loc)
	}
	{
		var location = 123
		var loc string
		loc = strconv.Itoa(location)

		fmt.Println(location, loc)
	}

	var m int64
	m = int64(marks)

	var i int32 = 100
	var r rune = i
	fmt.Println(r, i)

	fmt.Println(name, marks, address, m)
}
