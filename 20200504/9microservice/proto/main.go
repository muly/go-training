package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	c := Customer{Name: "C1", Address: "india"}

	b, err := proto.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
/////////////////////////////////////////////////////////////////////////////
	u := Customer{}
	err = proto.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)

}
