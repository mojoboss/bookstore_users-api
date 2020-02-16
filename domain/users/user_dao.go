package users

import (
	"fmt"
	"github.com/mojoboss/bookstore_users-api/datasources/postgres/users_db"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"log"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user &d not found", user.Id))
	}
	user.Id = result.Id
	user.Firstname = result.Firstname
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.insert_user($1, $2, $3)")
	if err != nil {
		log.Println("Error in db prepare for save user", err)
		return errors.NewInternalServerError("Server error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Firstname, user.LastName, user.Email)
	if err != nil {
		log.Println("Error in db exec for save user", err)
		return errors.NewInternalServerError("Server error")
	}
	if err != nil {
		log.Println("Error in getting last insert id for save user", err)
		return errors.NewInternalServerError("Server error")
	}
	return nil
}
