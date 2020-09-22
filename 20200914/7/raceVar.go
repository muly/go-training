package main

import("fmt")

var data string

func main() {
	c := make(chan bool)
	go func() {
		data = "a" // First conflicting access.
		c <- true
	}()
	data = "b" // Second conflicting access.
	<-c
	fmt.Println(data)
	
}