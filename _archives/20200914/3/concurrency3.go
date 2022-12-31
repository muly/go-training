// go routine() mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ilock sync.Mutex
var i int

func main() {

	fmt.Println("execution started")
	
	i=3
	
	wg.Add(1)
	go double()

	wg.Add(1)
	go double()

	// wg.Add(1)
	// go myPrint(3)

	// time.Sleep(2*time.Millisecond)
	wg.Wait()
	fmt.Println("execution completed, i=%v", i)
}

func double() {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Millisecond) // simulating some processing time
	// fmt.Println("myPrint", i)
	ilock.Lock() // lock
	i=i*2
	ilock.Unlock() // unlock
}


// func add(j int){
// 	i=i+j
// }