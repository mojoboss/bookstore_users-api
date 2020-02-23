package services

import (
	"github.com/mojoboss/bookstore_users-api/domain/users"
	"github.com/mojoboss/bookstore_users-api/utils/crypto_utils"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
)

var (
	UserService IUserService = &userService{}
)

type IUserService interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	SearchUser(string) (users.Users, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
}

type userService struct{}

func (service *userService) GetUser(userId int64) (user *users.User, err *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func (service *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.STATUS_ACTIVE
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (service *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	var user users.User
	return user.Search(status)
}

func (service *userService) DeleteUser(userid int64) *errors.RestErr {
	user, err := service.GetUser(userid)
	if err != nil {
		return err
	}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func (service *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := service.GetUser(user.Id)
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
		if user.Status != "" {
			current.Status = user.Status
		}
		if user.Password != "" {
			current.Password = user.Password
		}
	} else {
		current.Firstname = user.Firstname
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
		current.Password = user.Password
	}
	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
