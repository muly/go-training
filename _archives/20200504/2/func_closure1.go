package main

import (
	"fmt"
)

func main() {

	var a, b int = 1, 6
	su, av := func() (int, int) {
		sum := a + b
		avg := (a + b) / 2
		return sum, avg
	}()

	fmt.Println(su, av)
}
