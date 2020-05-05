package main

import (
	"fmt"
)

func main() {
	defer recoverFromPanic()
	defer test()
	divide(10,0)


	fmt.Println("program exited safely")
}


func divide(a, b int) (int) {
	return a / b
}

func recoverFromPanic(){
	if rec:= recover(); rec!= nil{
		fmt.Println("")
		fmt.Println("safely recovered from panic")
		return
	}
	fmt.Println("panic never happened")
}

func test(){
	fmt.Println("test executed")
}