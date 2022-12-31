package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

var mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	w.Write([]byte("invalid path" + r.URL.String()))
}

func main() {

	s := &http.Server{
		Addr:    ":8080",
		Handler: &myHandler{},
	}

	mux["/"] = helloHandler
	mux["/class"] = helloClass

	// r.HandleFunc("/", helloHandler)
	// r.HandleFunc("/class", helloClass)

	if err := s.ListenAndServe(); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! " + r.Method))
}

func helloClass(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Class ##########!"))
}
