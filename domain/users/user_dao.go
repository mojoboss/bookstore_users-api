package users

import (
	"fmt"
	"github.com/mojoboss/bookstore_users-api/datasources/postgres/users_db"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"log"
	"strings"
	"time"
)

const (
	DUPLICATE_INDEX_KEY = "users_email_key"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.get_user($1)")
	if err != nil {
		log.Println("Error in db prepare for get user", err)
		return errors.NewInternalServerError("Server error")
	}
	defer stmt.Close()
	var firstName, lastName, email string
	var creationTime time.Time
	err = stmt.QueryRow(user.Id).Scan(&firstName, &lastName, &email, &creationTime)
	if err != nil {
		log.Println("Error in scanning get user", err)
		return errors.NewInternalServerError("Server error")
	}
	user.Firstname = firstName
	user.LastName = lastName
	user.Email = email
	user.DateCreated = creationTime.String()
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
		if strings.Contains(err.Error(), DUPLICATE_INDEX_KEY) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalServerError("Server error")
	}
	if err != nil {
		log.Println("Error in getting last insert id for save user", err)
		return errors.NewInternalServerError("Server error")
	}
	return nil
}
