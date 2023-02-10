package endpoint

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type UserInformation struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	PaymentCard string `json:"payment_card"`
	Birthdate   string `json:"birthdate"`
	Phone       string `json:"phone"`
	Sex         string `json:"sex"`
}
