package main

import "fmt"

type studentMap map[string]int // user defined type

func (s studentMap) totalMarks() int {
	total := 0
	for _, v := range s {
		total = total + v
	}

	return total
}

func main() {
	var s studentMap
	s = map[string]int{"student1": 90, "student2": 50}
	
	fmt.Println(s.totalMarks())
}
