package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muly/go-training/20200504/9microservice/service"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCustomers API is called")
	vars := r.URL.Query()
	loc := vars["location"]
	if loc[0] == "" {
		http.Error(w, "missing location", http.StatusBadRequest)
		return
	}

	c := service.GetCustomers(loc[0])
	json.NewEncoder(w).Encode(c)
	w.Write()
}
