package main

import "fmt"

// for, break, continue

// for init, condition, inc {}

func main() {
	// // 5+2 iterations
	// for i:=1; i<=5; i++{
	// 	fmt.Println(i) 
	// }
	// for j:=1; j<=2; j++{
	// 	fmt.Println(j) 
	// }


	// // 5X2=10 iterations
	for i:=1; i<=5; i++{
		fmt.Println("i=",i) 
		for j:=1; j<=2; j++{
			fmt.Println("\tj=",j) 
		}
	}
	
}