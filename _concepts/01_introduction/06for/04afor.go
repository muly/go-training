package main

import "fmt"

func main() {
	// first n positive numbers whose sum is grater than 100
	sum := 0
	for i := 0; ; i++ {
		sum += i
		if sum > 100 {
			fmt.Printf("answer: first %d numbers\n", i)
			break
		}
	}
}
