package main

import "fmt"

func main() {
	var m map[string]string
	m = map[string]string{
		"123-456-7890": "John",
		"123-456-9999": "Max",
	}

	fmt.Println(m)
}
