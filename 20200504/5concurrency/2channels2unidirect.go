package main

import (
	"fmt"
	"time"
)

func main() {

	var r chan <- int = make(chan int) 
	var s <-chan int = make(chan int) 

	go ping(s,r)
	go pong(s)

	r <- 100

	time.Sleep(5 * time.Second)
}

func ping(s <- chan int, r chan <- int) {
	x:= <-s
	r<- x
}


func pong(c <- chan int) {
	x:= <- c
	fmt.Println(x)
}
