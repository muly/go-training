package main

import (
	"fmt"
)


func main() {
	// defer handlePanic()
	fmt.Println("started")
	n:=6
	n=f1()

	fmt.Println("finished", n)
}

func f1() int{
	defer handlePanic()
	var m map[string]int
	// m= map[string]int{} // commment out to simulate panic
	m["test"] = 0
	// panic("user generated panic")
	return 10
}


func handlePanic() {
	if r := recover(); r != nil { // recover from the panic here
		fmt.Println("recovered ", r)
		return
	}
	fmt.Println("no panic happened")
}
