package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go add(1, 5, &wg)

	wg.Add(1)
	go add(11, 5, &wg)

	wg.Wait() // time.Sleep(7 * time.Second)

}

func add(a, b int, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println(a + b)
	wg.Done()
}
