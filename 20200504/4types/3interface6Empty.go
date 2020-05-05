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
		fmt.Println(i)
}
