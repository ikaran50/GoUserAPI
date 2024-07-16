package handlers

import (
	"go-service/data"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func getIdFromURI(w http.ResponseWriter, r *http.Request) string {
	id := strings.TrimPrefix(r.URL.Path, "/")
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
	}
	return id
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.getUsers(rw)
		return
	case http.MethodPost:
		u.addNewUser(rw, r)
		return
	case http.MethodPut:
		idStr := getIdFromURI(rw, r)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		u.updateUsers(id, rw, r)
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

func (u *Users) updateUsers(id int, rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Update user")
	user := &data.User{}
	err := user.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to update user", http.StatusInternalServerError)
	}
	errMssg := data.UpdateUser(id, user)
	if errMssg != nil {
		http.Error(rw, "Unable to update user", http.StatusInternalServerError)
	}
}
