package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/muly/go-training/20200504/9microservice/api"
	// "github.com/muly/go-training/20200504/9microservice/data"
)

func main() {

	// err := data.Open()
	// if err != nil{
	// 	panic(err)
	// }

	h := mux.NewRouter()
	h.HandleFunc("/customer", api.GetCustomers).Methods("GET")
	http.ListenAndServe(":8080", h)
}
