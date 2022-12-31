// unit test

package main

import (
	"fmt"
	"time"
)

func main() {

	add(2, 0)

}

func add(a, b int) int {
	time.Sleep(5*time.Millisecond)
	for i:=0; i<10; i++{
		var j=i
		fmt.Println(j)
	}

	return a + b + 1
}
