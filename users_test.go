package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"github.com/gorilla/mux"
)

func TestGetUsers( t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")

	userJson := `{"firstname": "ardeshir", "lastname": "sepahsalar", "email": "ardeshir.org@gmail.com"}`
	req, err := http.NewRequest(
		"POST", 
		"/users",
		strings.NewReader(userJson),
	}
	if err != nil {
		t.Error(err)
	}
	
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
	}
}

func TestUniqueEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	
	userJson := `{"firstname": "ardeshir", "lastname": "sepahsalar", "email": "ardeshir.org@gmail.com"}`

	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Errorf("Bad Request expected: 400, got: %d", w.Code)
	}
}

func TestGetUsersClient( t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersUrl, nil)

	req, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got : %d", res.StatusCode)
	}
}

