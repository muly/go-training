package main

import (
	"fmt"
	"sync"
	"time"
)

var result int
var m sync.Mutex

func main() {
	go add(1, 5)
	go add(4, 5)
	go add(11, 5)
	go multiply(1, 5)

	time.Sleep(7 * time.Second)

	m.Lock()
	fmt.Println(result)
	m.Unlock()
}

func add(a, b int) {
	r := a + b
	time.Sleep(5 * time.Second)
	fmt.Println(a, b, r)

	m.Lock()
	result = r
	m.Unlock()
}

func multiply(a, b int) {
	r := a * b
	time.Sleep(5 * time.Second)
	fmt.Println(a, b, r)

	m.Lock()
	result = r
	m.Unlock()
}
