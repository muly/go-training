package main

import "fmt"

// Truth Table
// 				AND		OR
// 	FALSE FALSE FALSE	FALSE
// 	FALSE TRUE 	FALSE	TRUE
// 	TRUE FALSE 	FALSE	TRUE
// 	TRUE TRUE 	TRUE	TRUE

// && AND
// || OR
// !  NOT

func main() {
	// START OMIT
	var a, b int = 10, 2
	var class1, class2 = "go class", "java class"

	fmt.Println((a == b) && (class1 == class2))
	fmt.Println((a == b) || (class1 == class2))

	fmt.Println(!(a == b))

	fmt.Println(!((a == b) && (class1 == class2)))
	fmt.Println(!((a == b) || (class1 == class2)))
	// END OMIT
}
