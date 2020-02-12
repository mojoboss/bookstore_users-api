package users

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}
