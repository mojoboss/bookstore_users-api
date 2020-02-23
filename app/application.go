package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("starting the application..")
	router.Run(":8080")
}
