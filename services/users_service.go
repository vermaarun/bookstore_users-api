package services

import (
	"github.com/vermaarun/bookstore_users-api/domain/users"
	"github.com/vermaarun/bookstore_users-api/utils/crypto_utils"
	"github.com/vermaarun/bookstore_users-api/utils/date_time"
	"github.com/vermaarun/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreate = date_time.GetNowString()
	user.Status = users.UserStatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := users.User{Id: userId}
	return user.Delete()
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func GetAllUser() users.Users {
	return users.GetAll()
}

func Search(status string) (users.Users, *errors.RestError){
	dao := &users.User{}
	return dao.FindByStatus(status)
}