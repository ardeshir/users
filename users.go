package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var userStore = []User{}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)

}
