package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_users-api/domain/users"
	"github.com/mojoboss/bookstore_users-api/services"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: HANDLE THIS
		fmt.Println(err.Error())
		return
	}
	if err = json.Unmarshal(bytes, &user); err != nil {
		//TODO: HANDLE THIS
		fmt.Println(err.Error())
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: HANDLE THIS
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
