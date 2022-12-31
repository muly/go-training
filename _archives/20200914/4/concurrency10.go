
// channels for wait in place of wait group

package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup

func main() {
	fmt.Println("execution started")
	
	ch := make(chan int)

	// wg.Add(1)
	go myPrint(10, ch)


	go myPrint(10, ch)

	go myPrint(10, ch)


	// time.Sleep(2*time.Millisecond)
	// wg.Wait()
	_ = <- ch
	_ = <- ch
	_ = <- ch
	fmt.Println("execution completed")
}

func myPrint(i int, c chan int) {
	// defer wg.Done()
	time.Sleep(time.Duration(i) * time.Millisecond) // simulating some processing time
	fmt.Println("myPrint", i)

	c <- 1
}
