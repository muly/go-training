// channels for sync: buffered channel

package main

import (
	"fmt"
)
// var wg sync.WaitGroup

func main() {
	var s string
	var ch chan string = make(chan string, 2)



	s = "hi" 

	ch <- "hello1" // send
	ch <- "hello2" // send
	// ch <- "hello3" // send more than the buffer size is blockling operation

	s = <- ch // receive
	ch <- "hello3" // send is allowed as we emptied a slot in above step
	


	fmt.Println(s) // Hi or Hello ?
}

