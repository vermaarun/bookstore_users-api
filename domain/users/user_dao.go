package users

import "github.com/vermaarun/bookstore_users-api/utils/errors"

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError("User Not Found.")
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreate = result.DateCreate

	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("Email already registered.")
		}
		return errors.NewBadRequestError("User already Exist.")
	}
	usersDb[user.Id] = user
	return nil
}
