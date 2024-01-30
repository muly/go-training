package main

import "fmt"

func main() {

OuterLoop:
	for i := 1; i <= 5; i++ {
		fmt.Println("------------")
		for j := 1; j <= 10; j++ {
			if i*j > 15 {
				break OuterLoop
			}
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
