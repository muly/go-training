// channels for sync: buffered channel

package main

import (
	"fmt"
)

// var wg sync.WaitGroup

func main() {
	var ch chan int = make(chan int, 100)

	go func() {
		for i := 100; i < 300; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 300 // panic: send on closed channel
	}()

	for {
		j, ok:= <- ch
		if !ok{
			break
		}
		fmt.Println(j)
	}

	// time.Sleep(time.Minute)
}
