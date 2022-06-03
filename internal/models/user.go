package models

type UserModel struct {
	ID       uint    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
