package main

import (
	"fmt"
)

type myAdd func(a, b int) (int, int)

var f myAdd = func(a, b int) (int, int) {
	sum := a + b
	avg := (a + b) / 2
	return sum, avg
}

func main() {

	var f2 myAdd = func(a, b int) (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}

	{
		su, av := f(1, 6)
		fmt.Println(su, av)
	}
	{
		su, av := f2(10, 60)
		fmt.Println(su, av)
	}
}
