// channel select (withiut select)

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "a"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "b"
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}