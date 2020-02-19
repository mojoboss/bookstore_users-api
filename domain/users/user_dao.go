package users

import (
	"github.com/mojoboss/bookstore_users-api/datasources/postgres/users_db"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"github.com/mojoboss/bookstore_users-api/utils/postgres_utils"
	"log"
	"time"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.get_user($1)")
	if err != nil {
		log.Println("Error in db prepare for get user", err)
		return postgres_utils.HandlePQError(err)
	}
	defer stmt.Close()
	var firstName, lastName, email string
	var creationTime time.Time
	err = stmt.QueryRow(user.Id).Scan(&firstName, &lastName, &email, &creationTime)
	if err != nil {
		log.Println("Error in scanning get user", err)
		return postgres_utils.HandlePQError(err)
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
		return postgres_utils.HandlePQError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Firstname, user.LastName, user.Email)
	if err != nil {
		log.Println("Error in db exec for save user", err)
		return postgres_utils.HandlePQError(err)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.update_user($1, $2, $3, $4)")
	if err != nil {
		log.Println("Error in db prepare for save user", err)
		return postgres_utils.HandlePQError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Firstname, user.LastName, user.Email)
	if err != nil {
		log.Println("Error in db exec for save user", err)
		return postgres_utils.HandlePQError(err)
	}
	return nil
}
