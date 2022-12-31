// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r:= mux.NewRouter()
	r.HandleFunc("/api/", helloHandler)
	r.HandleFunc("/api/class", helloClass)
	r.Schemes("https")
	// http.Handle("/", r)



	   if err := http.ListenAndServeTLS(":443", "server.crt","server.key",r); err != nil{
	// if err := http.ListenAndServeTLS(":443", "server.crt","server.key",nil); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}


// helloHandler fdfdfdfdd
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! "+r.Method))
}

func helloClass(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Class!"))
}

//// to generate the cert files and key
// openssl genrsa -out server.key 2048
// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
