package main

import (
	"fmt"
	"time"
)

func main() {
	go add(1, 5)
	fmt.Println("called add function")

	time.Sleep(7 * time.Second)
}

func add(a, b int) {
	time.Sleep(5 * time.Second)
	fmt.Println(a + b)
}
