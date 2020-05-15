package users

import (
	"github.com/vermaarun/bookstore_users-api/datasources/mysql/users_db"
	"github.com/vermaarun/bookstore_users-api/utils/date_time"
	"github.com/vermaarun/bookstore_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name. email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users where id=?;"
	queryGetAllUser = "SELECT * FROM users;"
)

var (
	usersDb = make(map[int64]*User)
)

func GetAll() []User {

	stmt, err := users_db.Client.Prepare(queryGetAllUser)
	if err != nil {
		panic(err)
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	getResults, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	var userList []User
	var user = User{}

	// iterate over rows
	for getResults.Next() {
		// scan row
		err = getResults.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate)
		if err != nil {
			panic(err)
		}
		userList = append(userList, user)
	}

	//for _, user := range usersDb {
	//	userList = append(userList, User{
	//		Id:         user.Id,
	//		FirstName:  user.FirstName,
	//		LastName:   user.LastName,
	//		Email:      user.Email,
	//		DateCreate: user.DateCreate,
	//	})
	//}

	return userList
}

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	getResult := stmt.QueryRow(user.Id)
	if err = getResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(err.Error())
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// tell compiler to close statement before return
	defer stmt.Close()

	user.DateCreate = date_time.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.Id = userId
	return nil
}
