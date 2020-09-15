package main

import (
	"fmt"
)

func main() {

	// var m map[string]int
	// m := map[string]int{}
	m := make(map[string]int)
//  m["test"] = 100
	if m == nil{
		fmt.Println("empty")
	}else{
		fmt.Println("NOT empty")
	}
	
}