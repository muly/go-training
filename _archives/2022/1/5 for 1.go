package main

import "fmt"

// for, break, continue

// for init, condition, inc {}

func main() {
	for i:=0; i<10; i++{
		fmt.Println(i) // 0 to 9
	}

	j:=0
	for; j<10; j++{
		fmt.Println(j) // 0 to 9
	}

	for k:=0; k<10; {
		fmt.Println(k) // 0 to 9
		k++
	}

	l:=0
	for ; l<10; {
		fmt.Println(l) // 0 to 9
		l++
	}

	// infinite loop
	for{
		fmt.Println("hello") 
	}


}