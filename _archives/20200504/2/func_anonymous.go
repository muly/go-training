package main

import (
	"fmt"
)

func main() {
	su, av := func(a, b int) (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}(1, 6)

	fmt.Println(su, av)
}
