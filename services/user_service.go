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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.Firstname = user.Firstname
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
