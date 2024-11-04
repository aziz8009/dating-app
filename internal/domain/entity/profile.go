package entity

import "time"

type Profile struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
