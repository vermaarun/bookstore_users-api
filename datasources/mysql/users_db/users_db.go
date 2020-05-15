package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	MysqlUsersUsername = "mysql_users_username"
	MysqlUsersPassword = "mysql_users_password"
	MysqlUsersHost     = "mysql_users_host"
	MysqlUsersDb       = "mysql_users_db"
)

var (
	Client *sql.DB
	username = os.Getenv(MysqlUsersUsername)
	password = os.Getenv(MysqlUsersPassword)
	host = os.Getenv(MysqlUsersHost)
	db = os.Getenv(MysqlUsersDb)
)

func init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, db,
		)
	var err error
	Client, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured..!!")
}

