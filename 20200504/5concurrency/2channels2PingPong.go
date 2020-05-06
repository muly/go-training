package main

import (
	"fmt"
	"time"
)

func main() {

	var s chan int = make(chan int) 
	var r chan int = make(chan int) 

	go ping(s, r)
	go pong(s, r)

	r <- 100

	time.Sleep(5 * time.Second)
}


func ping(ping <-chan int, pong chan<- int) {
    for {
        <-ping
        fmt.Println("ping")
        time.Sleep(time.Second)
        pong <- 1
    }
}


func pong(ping chan<- int, pong <-chan int) {
    for {
        <-pong
        fmt.Println("pong")
        time.Sleep(time.Second)
        ping <- 1
    }
}