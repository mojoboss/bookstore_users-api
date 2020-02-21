package users

import (
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"strings"
)

const (
	STATUS_ACTIVE = "active"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("password field should be non-empty")
	}
	return nil
}
