package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go add(1, 5)

	go add(11, 5)

	wg.Wait() // time.Sleep(7 * time.Second)
}

func add(a, b int) {
	time.Sleep(5 * time.Second)
	fmt.Println(a + b)
	wg.Done()
}
