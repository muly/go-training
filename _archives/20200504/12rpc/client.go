package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

// Item : Non discript data type
type Item struct {
	Title string
	Body  string
}

var wg sync.WaitGroup

func main() {
	var reply1 Item
	var dB1 []Item
	// var reply2 Item
	// var dB2 []Item

	client1, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatalln("Connection error :", err)
	}
	defer client1.Close()

	// client2, err := rpc.DialHTTP("tcp", "localhost:4040")
	// if err != nil {
	// 	log.Fatalln("Connection error :", err)
	// }
	// defer client2.Close()

	a := Item{Title: "first", Body: "first item"}
	// b := Item{Title: "second", Body: "second item"}
	// c := Item{Title: "3rd", Body: "third test"}
	// d := Item{Title: "4th", Body: "4th test"}

	ch1 := make(chan *rpc.Call, 2)
	chw1 := client1.Go("API.AddItem", a, &reply1, ch1)
	fmt.Println("reply1", reply1)

	// // ch2:= make(chan *rpc.Call, 1)
	// chw2:= client2.Go("API.AddItem", b, &reply2, ch1)
	// fmt.Println("reply1", reply2)

	// client.Call("API.DeleteItem",c, &reply)
	// fmt.Println(reply)
	// fmt.Println()

	<-chw1.Done
	// <-chw2.Done

	fmt.Scanf("data inserted")

	// cha1:= make(chan *rpc.Call, 1)
	chaw1 := client1.Go("API.GetDB", "dB", &dB1, ch1)
	fmt.Println(dB1)

	// // cha2:= make(chan *rpc.Call, 1)
	// chaw2:=client2.Go("API.GetDB", "dB", &dB2, ch1)
	// fmt.Println(dB2)

	<-chaw1.Done
	//  <-chaw2.Done

	fmt.Println("wait done")

	fmt.Println(reply1)
	// fmt.Println(reply2)
	fmt.Println(dB1)
	// fmt.Println(dB2)
}
