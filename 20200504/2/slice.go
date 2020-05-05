//
package main

import "fmt"

func main() {

	// {
	// 	var a []int
	// 	a = []int{1, 2, 3, 4}
	// 	fmt.Println(a)
	// 	fmt.Println(len(a), cap(a))
	// }

	// {
	// 	var a [4]int
	// 	a = [4]int{1, 2, 3, 4}
	// 	fmt.Println(len(a), cap(a))

	// 	s := a[0:2]
	// 	s2 := a[2:]
	// 	fmt.Println(s2)
	// 	fmt.Println(len(s), cap(s))

	// 	// print(&a)
	// 	// fmt.Println()
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)

	// 	s = append(s, 10, 20)
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)

	// 	s = append(s, 30)
	// 	print(&s)
	// 	fmt.Println(" ", len(s), cap(s), s)
	// }

	// {
	// 	var a []int
	// 	var b []int
	// 	a = []int{1, 2, 3, 4}
	// 	b = a
	// 	fmt.Println(a)
	// 	fmt.Println(len(a), cap(a))
	// 	print(&a)

	// 	fmt.Println(b)
	// 	fmt.Println(len(b), cap(b))
	// 	print(&b)
	// 	fmt.Println()
	// }

	// {
	// 	a := []int{1, 2, 3, 4}
	// 	for _, b:= range a{
	// 		fmt.Println(b)
	// 	}
	// 	fmt.Println(len(a),cap(a))
	// }

	// {
	// 	var a []int
	// 	a = make([]int, 2, 4)

	// 	fmt.Println(len(a),cap(a))
	// 	fmt.Println(a)

	// 	a = append(a, 50)
	// 	fmt.Println(a)

	// 	a[1]= 60
	// 	fmt.Println(a)
	// }

	// {
	// 	var a = []int{1,2,3,4,5,6,7,8,9,0}
	// 	fmt.Println(a)
	// 	a= append(a[:4],a[5:]...)
	// 	fmt.Println(a, len(a), cap(a))
	// }

	// {
	// 	var a = []int{}
	// 	for i:=0; i<100; i++{
	// 		a=append(a,i)
	// 		fmt.Println(len(a), cap(a))
	// 	}
	// }

	{
		var a [4]int
		a = [4]int{1, 2, 3, 4}

		s := a[0:2]
		s[0] = 99
		fmt.Println(a)

		for i := 0; i < 5; i++ {
			s = append(s, i)
		}
		fmt.Println(s)
		fmt.Println(a)
	}

}
