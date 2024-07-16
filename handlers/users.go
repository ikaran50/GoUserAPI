package handlers

import (
	"go-service/data"
	"log"
	"net/http"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.getUsers(rw)
		return
	case http.MethodPost:
		u.addNewUser(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *Users) getUsers(rw http.ResponseWriter) {
	u.l.Println("Handle fetch Users")
	lu := data.GetUsers()
	err := lu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to create json", http.StatusInternalServerError)
	}
}

func (u *Users) addNewUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Create new user")
	user := &data.User{}
	err := user.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to create user", http.StatusInternalServerError)
	}
	data.AddUser(user)
}
