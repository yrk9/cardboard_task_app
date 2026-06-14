package model

import "time"

type User struct {
	ID int64
	Email string `json:"email"` 
	PasswordHash string	`json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}