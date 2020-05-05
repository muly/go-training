package main

import (
	"fmt"
)

func main() {
	f := func(a, b int) (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}
	{
		su, av := f(1, 6)
		fmt.Println(su, av)
	}
	{
		su, av := f(10, 60)
		fmt.Println(su, av)
	}
}
