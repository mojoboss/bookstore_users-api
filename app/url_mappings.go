package app

import "github.com/mojoboss/bookstore_users-api/controllers"

func mapUrls(){
	router.GET("/ping", controllers.Ping)
}