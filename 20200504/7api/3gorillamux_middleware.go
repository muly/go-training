package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type class struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {

	c := class{}

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods(http.MethodPost)
	r.Handle("/class/{classid}", MyMiddleware(c))
	r.HandleFunc("/class", postClass).Methods(http.MethodPost)

	// r.Use(MyMiddleware, MyMiddleware)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("restart serveer. encountered error, ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! " + r.Method))
}

func (class) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func MyMiddleware(h http.Handler) http.Handler {
	fmt.Println("running middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := strconv.Atoi(mux.Vars(r)["classid"]); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("given class id is not an number"))
			return
		}
		h.ServeHTTP(w, r)
	})
}
