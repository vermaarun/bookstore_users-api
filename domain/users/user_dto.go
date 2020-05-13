package users

import (
	"github.com/vermaarun/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id			int64 `json:"id"`
	FirstName	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreate 	string `json:"date_create"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}