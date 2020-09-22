//demonstrate example of http test using the http server code and the handler test (in _test.go)
package main

import (
	"net/http"
)

type message struct {
	Message string
}

func main() {
	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Hello World!"}`))
}
