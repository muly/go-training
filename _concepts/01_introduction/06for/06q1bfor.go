package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i >= 10 {
			continue
		}
		fmt.Println(i)
	}
}
