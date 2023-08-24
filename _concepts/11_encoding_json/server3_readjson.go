// basic hello world web application just using net/http package
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handler
//

var s1 student

type student struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"addr,omitempty"`
	Marks   int    `json:"marks,omitempty"`
}

func main() {
	// registring the handler to the specific URI
	http.HandleFunc("/student", handleStudent)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func handleStudent(w http.ResponseWriter, r *http.Request) {
	var s student

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			fmt.Println(err)
			// TODO
		}

		s1 = s

		fmt.Println(s)
	}else if r.Method == "GET"{
		err :=json.NewEncoder(w).Encode(s1)
		if err != nil {
			fmt.Println(err)
			// TODO
		}
	}

}
