package repository

import (
	"dating-app-backend/internal/domain/entity"
	"time"

	"gorm.io/gorm"
)

type SwipeRepository interface {
	Create(swipe *entity.Swipe) error
	CountTodaySwipes(userID uint) (int64, error)
}

type swipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{db}
}

func (r *swipeRepository) Create(swipe *entity.Swipe) error {
	return r.db.Create(swipe).Error
}

func (r *swipeRepository) CountTodaySwipes(userID uint) (int64, error) {
	var count int64
	todayStart := time.Now().Truncate(24 * time.Hour)
	if err := r.db.Model(&entity.Swipe{}).
		Where("user_id = ? AND created_at >= ?", userID, todayStart).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
