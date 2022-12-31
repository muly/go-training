package main

import (
	"net/http"
	"log"
	"time"

	"github.com/gorilla/mux"
)

func main() {


	r := mux.NewRouter()


	r.HandleFunc("/student/{id}", studentHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
	
}



func studentHandler (w http.ResponseWriter, r *http.Request){
	startTime := time.Now()
	vars := mux.Vars(r)

	id, ok :=vars["id"]
	if !ok{
		http.Error(w, "id missing", http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
	log.Printf("execution compelted. took: %v", startTime.Sub(time.Now()))
}