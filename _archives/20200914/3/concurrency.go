// go routine()
// channel type
// import "sync"
// // wait groups
// // mutex

// synchronization: wait groups, channel
// resource access: mutex, channel

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("execution started")

	wg.Add(1)
	go myPrint(10)

	wg.Add(1)
	go myPrint(2)

	wg.Add(1)
	go myPrint(3)

	// time.Sleep(2*time.Millisecond)
	wg.Wait()
	fmt.Println("execution completed")
}

func myPrint(i int) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Millisecond) // simulating some processing time
	fmt.Println("myPrint", i)
}
