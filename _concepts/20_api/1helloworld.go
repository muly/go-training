// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

// step 1: handler function definition
// when client sends a request to server endpoint, the request is finally routed to the corresponding handler function
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	// step 2: registering the handler function to the specific endpoint
	http.HandleFunc("/", hello)

	// step 3: port registration
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
