package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vermaarun/bookstore_users-api/domain/users"
	"github.com/vermaarun/bookstore_users-api/services"
	"github.com/vermaarun/bookstore_users-api/utils/errors"
	"net/http"
)

func GetUser(c *gin.Context) {}

func GetAllUser(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var user users.User

	// below commented code can be replaced by c.ShouldBindJSON() function
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// TODO: handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	// TODO: handle json error
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle json error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle save error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {}
