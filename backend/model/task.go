package model

import "time"

type Task struct {
	ID int64 `json:"id"`
	UserID int64 `json:"-"`
	Title string `json:"title"`
	Description string `json:"descriptions"`
	Priority string `json:"priority"`
	DueDate *time.Time `json:"due_date"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}