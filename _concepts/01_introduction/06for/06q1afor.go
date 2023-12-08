package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}
}
