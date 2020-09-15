package main

import (
	"fmt"
)

func main(){
	
	fmt.Println("started")
	f1()
	
	fmt.Println("finished")
}

func f1(){
	defer handlePanic()
	var m map[string]int
	//  m= map[string]int{} // commment out to simulate panic
	m["test"]=0

	// panic("user generated panic")
}


func handlePanic(){
	if r := recover(); r != nil { // recover from the panic here
		fmt.Println("recovered ", r)
		return
	}
	fmt.Println("no panic happened")
}