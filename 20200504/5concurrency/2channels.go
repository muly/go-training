package main

import (
	"fmt"
)

func main() {

	var c chan int // bidirectional channel

	c <- 5 // sending data into channel

	x:= <- c // receving from a channel

	fmt.Println(x)


	var receiveonly <- chan int 
	receiveonly <- 9 // invalid operation: receiveonly <- 9 (send to receive-only type <-chan int)

	var sendonly chan <- int 
	z :=  <- sendonly // invalid operation: <-sendonly (receive from send-only type chan<- int)

}

