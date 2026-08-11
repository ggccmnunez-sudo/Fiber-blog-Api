package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Views     int       `json:"views"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
