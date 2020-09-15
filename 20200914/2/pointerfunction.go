package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	fmt.Println(a, b) // 1 2
	swap(&a, &b)
	fmt.Println(a, b) // 2, 1
}

func swap(p, q *int){
	*p, *q  = *q, *p
}


