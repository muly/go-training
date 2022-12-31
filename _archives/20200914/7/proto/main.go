// protoc --go_out=. --go-grpc_out=.     --go_opt=paths=source_relative     --go-grpc_opt=paths=source_relative proto/customer.proto

package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	pb "github.com/muly/go-training/20200914/7/proto/proto"
)

func main() {

	c := pb.Customer{Name: "C1i", Address: "ndia"}

	b, err := proto.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b)) //C1india
	/////////////////////////////////////////////////////////////////////////////
	u := pb.Customer{}
	err = proto.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Address, u.Name)

}
