package models

import "time"

type Todo struct {
	ID          uint      `json:"id"`
	Task        string    `json:"task"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
