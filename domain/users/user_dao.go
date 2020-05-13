package users

import "github.com/vermaarun/bookstore_users-api/utils/errors"

var (
	usersDb = make(map[int64]*User)
)

func GetAll() []User {
	var userList []User
	for _, user := range usersDb{
		userList = append(userList, User{
			Id:         user.Id,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			DateCreate: user.DateCreate,
		})
	}
	return userList
}

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
