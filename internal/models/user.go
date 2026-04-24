package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
	Username string `json:"username"`
	Address  string `json:"address"`
}
