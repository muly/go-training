package main

import "fmt"

func main() {
	var num1, num2 int
	var class1, class2 = "go class", "java class"

	if (class1 == class2) && (num1 == num2) {
		fmt.Println("all are same")
	} else {
		fmt.Println("nothing is correct")
	}

}
