package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler(t *testing.T) {

	r, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.URL)
	//.Println(req.URL)

	w := httptest.NewRecorder()
	handler(w, r)

	t.Log(w.Body)
	t.Log(w.Code)

	if w.Code != http.StatusOK {
		t.Errorf("response test failed. wanted %v, but got %v", http.StatusOK, w.Code)
		return
	}

	m := struct{ Message string }{}

	json.NewDecoder(w.Body).Decode(&m)

	if m.Message != "hello" {
		t.Errorf("response test failed. wanted %v, but got %v", "hello", m.Message)
		return
	}

	//TODO: need to add reasponse body test
}
