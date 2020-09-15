package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	fmt.Println(a, b) // 1 2
	a, b= swap(a, b)
	fmt.Println(a, b) // 2 1
}

func swap(a, b int)(int, int){
	return b, a
}


