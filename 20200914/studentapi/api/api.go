package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/muly/go-training/20200914/studentapi/service"
)

func GetStudentHandler(session *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, ok := vars["id"]
		if !ok {
			http.Error(w, "id missing", http.StatusBadRequest)
			return
		}
		log.Printf("id: %s", id)

		s, err := service.GetStudent(session, id)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(s); err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

}
