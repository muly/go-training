package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/muly/go-training/20200914/studentapi/api"
)

var session *mgo.Session

func main() {

	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("error dial: ", err)
		return
	}

	r := mux.NewRouter()


	// POST
	// GET with filter; name 
	// DELETE
	// PUT
	r.HandleFunc("/student/{id}", api.GetStudentHandler(session)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))

}
