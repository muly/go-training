// channels for data sharing: as parameter to go routine

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("execution started")
	
	ch := make(chan int)

	// wg.Add(1)
	go add(10, 20, ch)


	go add(2,3, ch)

	go add(100, 400, ch)


	// time.Sleep(2*time.Millisecond)
	// wg.Wait()
	sum:= <- ch
	fmt.Println(sum )
	sum= <- ch
	fmt.Println( sum)
	sum= <- ch
	fmt.Println( sum)
	fmt.Println("execution completed")
}

func add(i , j int, c chan int) {
	time.Sleep(time.Duration(i) * time.Millisecond) // simulating some processing time
	c <- i+j
}
