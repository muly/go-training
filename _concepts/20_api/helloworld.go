// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// registering the handler to the specific URI
	http.HandleFunc("/", hello)

	// port registration
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// handler function definition
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
