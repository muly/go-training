package main

import (
	"fmt"
)

func main() {

	// // 
	// for init, check, inc
	// for i := 1; i <= 10; i++ {
	// 	fmt.Print(i)
	// }
	// fmt.Println()

	// // 
	// i := 1
	// for ; i <= 10; i++ {
	// 	fmt.Print(i)
	// }
	// fmt.Println()

	// // 
	// {
	// 	i := 1
	// 	for i <= 10 {
	// 		fmt.Print(i)
	// 		i++
	// 	}
	// 	fmt.Println()
	// }

	// 
	// {
	// 	i := 1
	// 	for {
	// 		fmt.Print(i)
	// 		if i >= 10 {
	// 			break
	// 		}
	// 		i++
	// 	}
	// 	fmt.Println()
	// }

	// {
	// 	i := 1
	// 	for {
	// 		if i >= 10 {
	// 			break
	// 		}
	// 		if i%2 == 0 {
	// 			i++
	// 			continue
	// 		}
	// 		fmt.Print(i)
	// 		i++
	// 	}
	// 	fmt.Println()
	// }

	// {
	// 	for j := 0; j < 3; j++ {
	// 		i := 1
	// 		for {
	// 			if i >= 10 {
	// 				break
	// 			}
	// 			if i%2 == 0 {
	// 				i++
	// 				continue
	// 			}
	// 			fmt.Print(i)
	// 			i++
	// 		}
	// 		fmt.Println()
	// 	}

	// 	fmt.Println()
	// }

Outerfor:
	{
		for j := 0; j < 3; j++ {
			i := 1
			for {
				if i >= 10 {
					break Outerfor
				}
				if i%2 == 0 {
					i++
					continue
				}
				fmt.Print(i)
				i++
			}
			fmt.Println()
		}

		fmt.Println()
	}
	


	s:= []int{2,4,6,8,10}
	for i, v:= range s{
		fmt.Println(i, v)
	} 
	for i, _:= range s{
		fmt.Println(i, s[i])
	} 
	for i := range s{
		fmt.Println(i, s[i])
	} 
	for _, v := range s{
		fmt.Println(v)
	} 
}
