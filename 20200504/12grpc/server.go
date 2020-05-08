package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	data := dataService{}

	s := grpc.NewServer()

	RegisterCustomerDataServer(s, data)

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln("listener error :", err)
	}

	log.Fatalln(s.Serve(listener))
}

type dataService struct{}

func (dataService) Insert(ctx context.Context, c *Customer) (*Customer, error) {
	fmt.Println(c.Name, c.Address)
	return &Customer{}, nil

}
func (dataService) GetAll(context.Context, *empty.Empty) (*Customers, error) {
	all := Customers{
		Records: []*Customer{
			&Customer{Name: "C222", Address: "usa"},
			&Customer{Name: "C333", Address: "uk"},
		},
	}
	return &all, nil
}
