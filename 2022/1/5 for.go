package main

import "fmt"

// for, break, continue

// for init, condition, inc {}

func main() {
	for i:=0; i<10; i++{
		
		fmt.Println(i) // 0 to 9
	}
	for i:=10; i>0; i--{
		fmt.Println(i) // 10 to 1
	}
}