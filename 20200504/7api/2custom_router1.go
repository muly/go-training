// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

func main() {

	r:= http.NewServeMux()
	
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/class", helloClass)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("Hello World! "+r.Method))
}

func helloClass(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Class!"))
}
