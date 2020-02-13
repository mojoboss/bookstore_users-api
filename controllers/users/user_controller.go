package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_users-api/domain/users"
	"github.com/mojoboss/bookstore_users-api/services"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO:HANDLE ERROR return bad request
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: HANDLE THIS user creation error
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "later")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "later")
}
