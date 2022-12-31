// channels for sync

package main

import (
	"fmt"
)
// var wg sync.WaitGroup

func main() {
	var s string
	var ch chan string = make(chan string)



	s = "hi" 
	go func(){
		ch <- "hello" // send
	}()
	s = <- ch // receive
	
	//// OR
	// go func(){
	// 	s = <- ch // receive
	// }()
	// ch <- "hello" // send
	

	fmt.Println(s) // Hi or Hello ?
}

