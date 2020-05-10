package api

import (
	"encoding/json"
	"net/http"

	"github.com/muly/go-training/20200504/9microservice/service"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	loc := vars["location"]
	if loc[0] == "" {
		http.Error(w, "missing location", http.StatusBadRequest)
		return
	}


	defer r.Body.Close()

	report, err := service.GetReports(loc[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(report)
}
