// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

type class struct {
}

func (class) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Class!"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! " + r.Method))
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.Handle("/class", class{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}
