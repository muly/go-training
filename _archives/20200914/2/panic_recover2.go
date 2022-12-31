package main

import (
	"fmt"
)

func main() {
	// defer handlePanic()
	fmt.Println("started")
	f4()
	fA()

	fmt.Println("finished")
}

func f4() {	
	defer handlePanic()
	f3()
}

func f3() {
	f2()
}

func f2() {
	f1()

}

func f1() {
	var m map[string]int
	// m= map[string]int{} // commment out to simulate panic
	m["test"] = 0
	// panic("user generated panic")
}

func fA() {
	defer handlePanic()
	panic("user generated panic")
}

func handlePanic() {
	if r := recover(); r != nil { // recover from the panic here
		fmt.Println("recovered ", r)
		return
	}
	fmt.Println("no panic happened")
}
