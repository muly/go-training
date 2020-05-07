package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan int)
	m := make(chan int)

	go add(1, 5, a)
	go multiply(11, 5, m)

	// out := <- a
	// fmt.Println(out)
	// out = <- m
	// fmt.Println(out)

	out := 0

	select {
	case out = <-a:
		fmt.Println(out)
	case out = <-m:
		fmt.Println(out)
	default:
		time.Sleep(4 * time.Second) // TODO: not working
	}

}

func add(a, b int, out chan int) {
	time.Sleep(5 * time.Second)
	out <- (a + b)
}

func multiply(a, b int, out chan int) {
	time.Sleep(1 * time.Second)
	out <- (a * b)
}
