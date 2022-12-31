// endpoints: /hello [GET] -> "hello class"

package main

import (
	"fmt"
	"net/http"
)

var data []student

func main() {
	// host (ip address): localhost 127.0.0.1
	// port: 8080
	// endpoint: /hello
	// method: GET

	// registration: route, 	handler function
	// registration: route, 	handler

// /hel
	http.HandleFunc("/student", studentHandler)
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/logout", logoutHandler)
	// http.Handle()
	

	// listen on port
	http.ListenAndServe("localhost:8082", nil)
}



 func studentHandler (w http.ResponseWriter, r *http.Request){

	switch r.Method{
	case http.MethodGet:
		http.Write
	}


	// w.Write()

fmt.Fprintf(w, 	"hello class")
}