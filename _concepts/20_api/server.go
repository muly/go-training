// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

// handler registration
// port registration
// handler function definition

func main() {
	// registering the handler to the specific URI
	http.HandleFunc("/student", create)
	http.HandleFunc("/address", hello)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	w.WriteHeader(400)
	w.Write([]byte("Create World!"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
