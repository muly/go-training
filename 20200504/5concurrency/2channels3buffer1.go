package main

import (
	"fmt"
)

func main() {

	var c chan int = make(chan int, 1) // bidirectional channel

	c <- 5 // sending data into channel
	c <- 15


	fmt.Println(x)

}