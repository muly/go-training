package main

import ("fmt")


type dog struct {
	name  string
	owner string
	address string
}


func main(){
	var i int = 10
	myFunc(i)

	var s string = "hello"
	myFunc(s)

	var d dog = dog{"max", "owner1", "india"}
	myFunc(d)

	var i64 int64 = 100
	myFunc(i64)
}


func myFunc(i interface{}){
	switch i.(type){
	case int:
		fmt.Printf("found int %v \n", i)
	case string:
		fmt.Printf("found string %v \n", i)
	case dog:
		fmt.Printf("found dog %v \n", i)
	default:
		fmt.Printf("found UNKNOWN type %v \n", i)
	}

}
