package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan int = make(chan int) // bidirectional channel

	go double(c)
	c <- 5 // sending data into channel

	x := <-c
	fmt.Println(x)

}

func double(c chan int) {
	x := <-c
	time.Sleep(5 * time.Second)
	c <- x * 2
}
