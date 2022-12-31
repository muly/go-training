// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", helloHandler)
	http.HandleFunc("/api/class", helloClass)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}


// helloHandler fdfdfdfdd
func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.:

	case "POST":
		r.Body
		r.Header
	}
	w.Write([]byte("Hello World! "+r.Method))
}

func helloClass(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Class!"))
}
