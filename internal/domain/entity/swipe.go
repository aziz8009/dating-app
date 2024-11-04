package entity

import "time"

type Swipe struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProfileID uint      `json:"profile_id"`
	Action    string    `json:"action"` // "like" or "pass"
	CreatedAt time.Time `json:"created_at"`
}
