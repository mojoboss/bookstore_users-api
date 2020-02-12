package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_users-api/domain/users"
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
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "later")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "later")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "later")
}
