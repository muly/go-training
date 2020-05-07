package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/muly/go-training/20200504/9microservice/api"
)

func main() {
	h := mux.NewRouter()
	h.HandleFunc("/report", api.GetReports).Methods("GET")
	http.ListenAndServe(":8080", h)
}
