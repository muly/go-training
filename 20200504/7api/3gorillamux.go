package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type class struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods(http.MethodPost)
	r.HandleFunc("/class/{classid}", helloClass)
	r.HandleFunc("/class", postClass).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! " + r.Method))
}

func helloClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte("Hello Class, your id is " + vars["classid"] + " with query " + vars["search"]))
}

func postClass(w http.ResponseWriter, r *http.Request) {

	c := class{}

	json.NewDecoder(r.Body).Decode(&c)
	defer r.Body.Close()

	c.Name = "hello " + c.Name

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(c); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

}
