package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Users is a collection of user
type Users []*User

var errorMessage = fmt.Errorf("User Not found")

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// GetUsers returns a list of users
func GetUsers() Users {
	return userList
}

func AddUser(u *User) {
	latestId := userList[len(userList)-1].ID
	u.ID = latestId + 1
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	foundUser, pos, err := FindExistingUser(id)
	if foundUser != nil {
		userList[pos] = u
		return nil
	}
	return err
}

func FindExistingUser(id int) (*User, int, error) {
	for pos, u := range userList {
		if u.ID == id {
			return u, pos, nil
		}
	}
	return nil, -1, errorMessage
}

var userList = []*User{
	&User{
		ID:        1,
		Name:      "Ishaan",
		Location:  "Canada",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&User{
		ID:        2,
		Name:      "Joe",
		Location:  "US",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
