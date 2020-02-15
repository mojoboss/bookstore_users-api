package services

import (
	"github.com/mojoboss/bookstore_users-api/domain/users"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (user *users.User, err *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
