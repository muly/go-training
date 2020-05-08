package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:5050", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := NewCustomerDataClient(conn)

	ctx := context.TODO()

	a := Customer{Name: "C111", Address: "india"}
	_, err = c.Insert(ctx, &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	all, err := c.GetAll(ctx, &empty.Empty{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("received:", all.Records)
}
