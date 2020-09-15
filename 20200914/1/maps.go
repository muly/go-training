package main

import (
	"fmt"
)

func main() {

	// var m map[string]int
	// m := map[string]int{}
	m := make(map[string]int)

	m["student1"]=120

	_, ok := m["student1"]
	if ok{
		fmt.Println("key exists")
	}else{
		fmt.Println("key doesn't exists")
	}

	delete(m, "student1")

	_, ok = m["student1"]
	if ok{
		fmt.Println("key exists")
	}else{
		fmt.Println("key doesn't exists")
	}

	// for k, v:=range m{
	// 	fmt.Println(k, v)
	// }
}
