package repository

import (
	"dating-app-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if user is not found
		}
		return nil, err
	}
	return &user, nil
}
