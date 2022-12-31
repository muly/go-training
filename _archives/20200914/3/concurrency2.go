// go routine() pointer parameter

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("execution started")
	var i *int = new(int) 
	*i = 10

	// *i := 10

	wg.Add(1)
	go double(i)

	// wg.Add(1)
	// go myPrint(2)

	// wg.Add(1)
	// go myPrint(3)

	// time.Sleep(2*time.Millisecond)
	wg.Wait()
	fmt.Println("execution completed, i=%v", *i)
}

func double(i *int) {
	defer wg.Done()
	time.Sleep(time.Duration(*i) * time.Millisecond) // simulating some processing time
	// fmt.Println("myPrint", i)
	*i=*i*2
}
