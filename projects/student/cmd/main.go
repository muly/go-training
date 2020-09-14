package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/muly/go-training/projects/student/api"
)

func main() {

	port := "8080"

	router := mux.NewRouter()

	router.HandleFunc("/students", jsonHeaders(api.HandlePost)).Methods("POST")
	router.HandleFunc("/students", jsonHeaders(api.HandleGet)).Methods("GET")
	router.HandleFunc("/students/{id}", jsonHeaders(api.HandleGetById)).Methods("GET")
	router.HandleFunc("/students/{id}", jsonHeaders(api.HandlePut)).Methods("PUT")
	router.HandleFunc("/students/{id}", jsonHeaders(api.HandlePatch)).Methods("PATCH")
	router.HandleFunc("/students/{id}", jsonHeaders(api.HandleDelete)).Methods("DELETE")

	log.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println("Listening error :", err)
	}
}

// jsonHeaders sets header with content type as application/json for all pages
func jsonHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}
