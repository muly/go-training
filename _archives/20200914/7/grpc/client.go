package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	customerpb "github.com/muly/go-training/20200914/7/grpc/proto"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := customerpb.NewCustomerDataClient(conn)

	ctx := context.TODO()

	a := customerpb.Customer{Name: "C111", Address: "india"}
	_, err = c.Insert(ctx, &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	a = customerpb.Customer{Name: "C222", Address: "usa"}
	_, err = c.Insert(ctx, &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	a = customerpb.Customer{Name: "C333", Address: "uk"}
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

	fmt.Println("received:", all.Data)
}
