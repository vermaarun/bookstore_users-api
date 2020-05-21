package users

import "encoding/json"

type PublicUser struct {
	Id			int64 `json:"id"`
	DateCreate 	string `json:"date_create"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id			int64 `json:"id"`
	FirstName	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreate 	string `json:"date_create"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{}{
	// 1. One way of creating user
	if isPublic {
		return PublicUser{
			Id: user.Id,
			DateCreate: user.DateCreate,
			Status: user.Status,
		}
	}

	// 1. Second way of creating user
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}