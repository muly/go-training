package main

// protoc --go_out=. --go-grpc_out=.     --go_opt=paths=source_relative     --go-grpc_opt=paths=source_relative proto/customer.proto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	customerpb "github.com/muly/go-training/20200914/7/grpc/proto"
)

func main() {

	implementation := customerpb.CustomerDataService{}
	implementation.Insert = Insert
	implementation.GetAll = GetAll

	s := grpc.NewServer()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	customerpb.RegisterCustomerDataService(s, &implementation)

	log.Fatal(s.Serve(l))

}

var data []customerpb.Customer

func Insert(ctx context.Context, c *customerpb.Customer) (*customerpb.Customer, error) {
	data = append(data, *c)
	return c, nil
}

func GetAll(context.Context, *empty.Empty) (*customerpb.Customers, error) {
	out := customerpb.Customers{}
	for _, v := range data {
		out.Data = append(out.Data, &v)
	}

	return &out, nil
}
