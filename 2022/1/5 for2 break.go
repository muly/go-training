package main

import "fmt"

// for, break, continue

func main() {
	for i := 0;i<=10 ; i = i + 1 {
		fmt.Println(i)
	}
	
	for i := 0; ; i = i + 1 {
		fmt.Println(i)
		if i > 10 {
			break
		}
		fmt.Println("hello")
	}
}
