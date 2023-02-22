package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	EmailAddress string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}