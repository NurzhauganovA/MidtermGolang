package endpoint

import "time"

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInformation struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	PaymentCard string    `json:"payment_card"`
	Birthdate   time.Time `json:"birthdate"`
	Phone       string    `json:"phone"`
	Sex         string    `json:"sex"`
}
