// sync between multiple go routines: relay race simulation

package main

import (
	"fmt"
	"time"
)



func main(){
	fmt.Println("execution started")
	
	ch := make(chan int)
	done := make(chan int)
	// 
		go pass(ch,done)
		go pass(ch,done)
		go pass(ch,done)
		go pass(ch,done)

		ch <- 1 // start the race

		_= <- done

	// _ = <-ch
	// _ = <-ch2

	fmt.Println("execution completed")
}

func pass(c chan int, done chan int) {
	i:= <- c
	time.Sleep(time.Duration(i) * time.Second) // simulating some processing time
	fmt.Println("running")
	if i==4{
		done <-1
		return
	}
	c <- i+1
}
