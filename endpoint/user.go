package endpoint

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"firstname"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
