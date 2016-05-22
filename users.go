package users

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email	  string `json:"email"`
}

