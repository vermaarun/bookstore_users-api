package users

import (
	"fmt"
	"github.com/vermaarun/bookstore_users-api/datasources/mysql/users_db"
	"github.com/vermaarun/bookstore_users-api/logger"
	"github.com/vermaarun/bookstore_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows                 = "no rows in result set"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=?;"
	queryGetAllUser             = "SELECT id, first_name, last_name, email, date_created, status FROM users;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser             = "DELETE FROM users where id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users where status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users where email=? AND password=? AND status=?;"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var userList []User

	// iterate over rows
	for rows.Next() {
		// scan row
		var user User
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate, &user.Status)
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		userList = append(userList, user)
	}
	if len(userList) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user matches status %s", status))
	}
	return userList, nil
}

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
	defer getResults.Close()
	var userList []User
	var user = User{}

	// iterate over rows
	for getResults.Next() {
		// scan row
		err = getResults.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate, &user.Status)
		if err != nil {
			panic(err)
		}
		userList = append(userList, user)
	}

	return userList
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	getResult := stmt.QueryRow(user.Id)
	if err = getResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate, &user.Status); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(err.Error())
		}
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save statement", err)
		return errors.NewInternalServerError("database error")
	}
	// tell compiler to close statement before return
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreate, user.Status, user.Password)
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

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) FindByEmailAndPassword() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	getResult := stmt.QueryRow(user.Email, user.Password, UserStatusActive)
	if err = getResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate, &user.Status); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("Invalid user credentials")
		}
		logger.Error("error when trying to find user by email and password", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
