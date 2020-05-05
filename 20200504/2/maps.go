//
package main

import "fmt"

func main() {
	// {
	// 	var m map[int]string

	// 	if m ==nil{
	// 		fmt.Println("map in nil")
	// 	}

	// 	m = make(map[int]string)

	// 	if m ==nil{
	// 		fmt.Println("map in nil")
	// 	}

	// 	m[0] = "A"
	// 	m[9] = ""
	// 	fmt.Println(m[0])

	// 	m[0] = "B"
	// 	fmt.Println(m[0])

	// 	{
	// 		v, ok := m[9]
	// 		fmt.Println(v, ok)
	// 	}

	// 	{
	// 		v, ok := m[19]
	// 		fmt.Println(v, ok)
	// 	}

	// 	{
	// 		k := 0
	// 		fmt.Println(m[k])
	// 	}
	// }

	// {
	// 	var m map[int]*string
	// 	m = make(map[int]*string)
	// 	fmt.Println(m[0])
	// 	{
	// 		v, ok := m[0]
	// 		fmt.Println(v, ok)
	// 	}

	// 	m[9] = nil
	// 	{
	// 		v, ok := m[9]
	// 		fmt.Println(v, ok)
	// 	}

	// }

	// {
	// 	var m map[int]string
	// 	m = map[int]string{} // m = make(map[int]string)
	// 	fmt.Println(m)
	// }

	// {
	// 	var m map[int]string
	// 	m = map[int]string{0: "A", 9: "B"}
	// 	fmt.Println(m, len(m))

	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// 	for _, v := range m {
	// 		fmt.Println(v)
	// 	}

	// 	for k, _ := range m {
	// 		fmt.Println(k)
	// 	}
	// 	for k := range m {
	// 		fmt.Println(k)
	// 	}
	// }

	{
		var m map[int]string
		m = map[int]string{0: "A", 9: "B"}
		fmt.Println(m[9])
		m[9] = ""
		{
			_, ok := m[9]
			fmt.Println(m[9], ok)
		}
		delete(m, 9)
		{
			_, ok := m[9]
			fmt.Println(m[9], ok)
		}
	}

}
