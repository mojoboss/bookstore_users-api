package users

import (
	"github.com/mojoboss/bookstore_users-api/datasources/postgres/users_db"
	"github.com/mojoboss/bookstore_users-api/logger"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"time"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.get_user($1)")
	if err != nil {
		logger.Error("Error in Prepare get_user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	var firstName, lastName, email, status, password string
	var creationTime time.Time
	err = stmt.QueryRow(user.Id).Scan(&firstName, &lastName, &email, &status, &password, &creationTime)
	if err != nil {
		logger.Error("Error in Query row in get_user", err)
		return errors.NewInternalServerError("database error")
	}
	user.Firstname = firstName
	user.LastName = lastName
	user.Email = email
	user.Status = status
	user.Password = password
	user.DateCreated = creationTime.String()
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.insert_user($1, $2, $3, $4, $5)")
	if err != nil {
		logger.Error("Error in Prepare insert_user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Firstname, user.LastName, user.Email, user.Status, user.Password)
	if err != nil {
		logger.Error("Error in exec for insert_user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.update_user($1, $2, $3, $4, $5, $6)")
	if err != nil {
		logger.Error("Error in db prepare for update_user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Firstname, user.LastName, user.Email, user.Status, user.Password)
	if err != nil {
		logger.Error("Error in db exec for update_user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.search_user($1)")
	if err != nil {
		logger.Error("Error in search_user prepare", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error in search_user query", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Firstname, &usr.LastName, &usr.Email, &usr.Status, &usr.DateCreated); err != nil {
			logger.Error("Error in scan of search_user", err)
			return nil, errors.NewInternalServerError("database error")
		}
		users = append(users, usr)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError("No users found for the status")
	}
	return users, nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("SELECT * FROM users_db.delete_user($1)")
	if err != nil {
		logger.Error("Error in db prepare for delete_user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("Error in db exec for delete_user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
