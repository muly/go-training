// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

// handler registration
// port registration
// handler function definition

// step 1: handler function definitions
func student(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	if r.Method == "POST" {
		// student record created in db
		w.WriteHeader(201)
		w.Write([]byte("student created"))
	} else if r.Method == "GET" {
		// w.WriteHeader(200) // 200 is default
		w.Write([]byte("here is the student information"))
	} else {
		w.WriteHeader(501)
	}
}
func address(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
}

func main() {
	// step 2: registering the handler function to the specific endpoint
	http.HandleFunc("/student", student)
	http.HandleFunc("/address", address)

	// step 3: port registration
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
