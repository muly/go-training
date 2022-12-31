// type switching

package main

import "fmt"

type math interface {}

func main() {

	var i int = 10
	process(i)


	var s string = "100"
	process(s)


	var f float32 = 99.99
	process(f)
}

func process(m math) {
	switch m.(type){
	case int:
		fmt.Printf("%v: int type\n", m)
	case string:
		fmt.Printf("%v: string type\n", m)
	default:
		fmt.Printf("%v: unknown type\n", m)
	}

	switch t := m.(type){
	case int:
		fmt.Printf("%d: int type\n", t)
	case string:
		fmt.Printf("%v: string type\n", m)
	default:
		fmt.Printf("%v: unknown type\n", m)
	}

}
