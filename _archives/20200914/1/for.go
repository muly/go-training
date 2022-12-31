package main

import("fmt")

// _, v
// i, v
func main(){
	s := []int{2,4,5,6}
	
	// for i:=0; i<len(a); i++{
	for _, v:= range s{
		if v % 2 == 0{
			continue
		}
		fmt.Println(v)	
	}
}