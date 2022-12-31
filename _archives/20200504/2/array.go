//
package main

import "fmt"

func main() {

	{
		var a [4]int
		a = [4]int{1, 2, 3, 4}
		fmt.Println(a)

		fmt.Println(a[0:2])
		// fmt.Println(a[4])
	}

	// {
	// 	var a = [4]int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// {
	// 	a := [4]int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// }

	// {
	// 	a := [4]int{1, 2, 3, 4}
	// 	var b [4]int
	// 	b = a
	// 	fmt.Println(a, b)
	// }

	// {
	// 	a := [4]int{1, 2, 3, 4}
	// 	var b [5]int
	// 	b = a
	// 	fmt.Println(a, b)
	// }

	// {
	// 	var a [4]int
	// 	fmt.Println(a)
	// }

	// {
	// 	a := [4]int{1, 2, 3, 4}
	// 	for _, b:= range a{
	// 		fmt.Println(b)
	// 	}
	// 	fmt.Println(len(a),cap(a))
	// }

}
