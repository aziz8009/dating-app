package repository

import (
	"dating-app-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(profile *entity.Profile) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}

func (r *profileRepository) Create(profile *entity.Profile) error {
	return r.db.Create(profile).Error
}
