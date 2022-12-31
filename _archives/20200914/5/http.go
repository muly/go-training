// endpoints: /hello [GET] -> "hello class"

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// host (ip address): localhost 127.0.0.1
	// port: 8080
	// endpoint: /hello
	// method: GET

	// registration: route, 	handler function
	// registration: route, 	handler

	// /hel
	http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/logout", logoutHandler)
	// http.Handle()

	// listen on port
	http.ListenAndServe(":8082", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello class"))
	// OR
	fmt.Fprintf(w, "hello class")
}

/*
http.ResponseWriter interface{
	Write([]byte) (int, error)
	also other methods
}

io.Writer interface{
	Write([]byte) (int, error)
}

if w implements http.ResponseWriter, then it also implements io.Writer

	type s struct{}
	func (s)Write(){	}
	func (s)Something(){	}

	s1 := s{}

	var rw http.ResponseWriter
	rw = s1 // ERROR

	var w2 io.Writer
	w2 = s1 /// ALLOWED

	type t struct{}
	func (t)Write(){	}
	func (t)also other methods(){	}

	t1 := t{}
	rw = t1 // allowed
	w2 = t1

	fmt.Fprintf(w)

*/
