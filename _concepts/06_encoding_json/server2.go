// basic hello world web application just using net/http package
package main

import (
	"net/http"
)

// handler
//

func main() {
	// registring the handler to the specific URI
	http.HandleFunc("/student", hello)
	http.ListenAndServe(":9090", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello New World!"))
}
