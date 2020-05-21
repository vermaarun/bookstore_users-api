package users

import (
	"github.com/vermaarun/bookstore_users-api/utils/errors"
	"strings"
)

const (
	UserStatusActive = "active"
)

type User struct {
	Id			int64 `json:"id"`
	FirstName	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreate 	string `json:"date_create"`
	Status      string `json:"status"`
	Password	string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}