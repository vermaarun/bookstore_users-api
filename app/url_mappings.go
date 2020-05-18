package app

import "github.com/vermaarun/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users", controllers.GetAllUser)
	//router.GET("/users/search", controllers.FindUser)

	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)

	router.DELETE("/users/:user_id", controllers.DeleteUser)
}