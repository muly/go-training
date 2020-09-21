package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var mw = alice.New(logger)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("execution compelted. took: %v", time.Since(startTime))
	})

}

func main() {

	r := mux.NewRouter()

	r.Handle("/student/{id}", mw.ThenFunc(studentHandler)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))

}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		http.Error(w, "id missing", http.StatusBadRequest)
		return
	}

	time.Sleep(time.Second)

	w.Write([]byte(id))
	
}
