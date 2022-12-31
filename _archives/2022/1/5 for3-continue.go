package main

import "fmt"

// for, break, continue

func main() {
	for i := 1; i<5; i = i + 1 {
		fmt.Println("hello", i)

		if i == 2{
			continue // skips the rest of the iteration and go to the next iteration
			// break // skips the rest of the for loop and go to the next command after the for loop
		}

		fmt.Println("world", i)

	}

}
