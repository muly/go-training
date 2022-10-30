package main

import "fmt"

// for, break, continue

func main() {

	for i := 1; i<100; i = i + 1 {
		if i%10==0{
			continue
		}

		if i%2==0 {
			fmt.Println(i, " even ")
		}

	}
	fmt.Println("end")
}
