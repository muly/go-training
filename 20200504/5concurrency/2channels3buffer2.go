package main

import (
	"fmt"
)

func main() {

	var c chan int = make(chan int, 2) // bidirectional channel

	c <- 5 // sending data into channel
	c <- 15

	x := <-c // receving from a channel
	y := <-c

	fmt.Println(x, y)

}
