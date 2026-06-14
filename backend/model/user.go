package model

import "time"

type User struct {
	ID int
	Email string `json:"email"` 
	PasswordHash string	`json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}