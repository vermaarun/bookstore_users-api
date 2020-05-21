package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vermaarun/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("starting application...")
	router.Run(":8080")
}
