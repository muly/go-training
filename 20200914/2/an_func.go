package main

import (
	"fmt"
)


func main() {
	a, b := 1, 2
	fmt.Println(a, b) // 1 2
	
	// func(p, q *int){
	// 	*p, *q  = *q, *p
	// }(&a, &b)

	// func(){
	// 	b, a  = a, b
	// }()

	// var myfunc = func(){
	// 	b, a  = a, b
	// }
	// myfunc()
	
	fmt.Println(a, b) // 2, 1
}







