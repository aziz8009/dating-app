package entity

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-"`
	Premium   bool      `json:"premium"`
	CreatedAt time.Time `json:"created_at"`
}
