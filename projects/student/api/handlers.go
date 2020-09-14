package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/muly/go-training/projects/student/crud"
	data "github.com/muly/go-training/projects/student/db"

	"github.com/gorilla/mux"
)

// HandlePost saves the data in the db.
func HandlePost(w http.ResponseWriter, r *http.Request) {

	students := []data.Student{}

	if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := crud.Post(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandleGet get the filtered data.
func HandleGet(w http.ResponseWriter, r *http.Request) {
	// TODO: filters
	// params, err := url.ParseQuery(r.URL.RawQuery)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	reuslts, err := crud.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(reuslts) == 0 {
		fmt.Fprintf(w, "There are no data for such query")
		return
	}

	if err := json.NewEncoder(w).Encode(reuslts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandleGetById get the data by name of the student provided.
func HandleGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	id := params["id"]

	results, err := crud.GetById(id)
	if err != nil {
		http.Error(w, err.Error()+": "+id, http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandlePut gets the data by name provided and replaces the content.
func HandlePut(w http.ResponseWriter, r *http.Request) {
	results := data.Student{}

	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&results); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := crud.Put(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	note, err := crud.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandlePatch updates only the given fields from the request body.
// TODO: key fields are not allowed to be updated
func HandlePatch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "problem to read request body", http.StatusBadRequest)
	}

	data := map[string]interface{}{}
	if err = json.Unmarshal(content, &data); err != nil {
		http.Error(w, fmt.Sprintf("Unmarshal error.. %s: %v", content, err), http.StatusBadRequest)
		return
	}

	student, err := crud.Patch(id, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandleDelete identifies by name provided and deletes it from db.
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := crud.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Record deleted successfully")
}
