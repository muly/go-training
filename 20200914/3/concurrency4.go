// channels


//// send, receive
//// blocking operation: send, receive
//// 


package main

import (
	"fmt"
)
// var wg sync.WaitGroup

func main() {
	var s string
	var ch chan string = make(chan string)

	ch <- "hello" // send

	s = "hi" 
	s = <- ch // receive

	fmt.Println(s)
}

