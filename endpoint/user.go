package endpoint

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"firstname" binding:"required"`
	Surname  string `json:"surname"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
