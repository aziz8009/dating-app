package repository

import (
	"dating-app-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type PremiumRepository interface {
	Create(pkg *entity.PremiumPackage) error
	GetAll() ([]entity.PremiumPackage, error)
}

type premiumRepository struct {
	db *gorm.DB
}

func NewPremiumRepository(db *gorm.DB) PremiumRepository {
	return &premiumRepository{db}
}

func (r *premiumRepository) Create(pkg *entity.PremiumPackage) error {
	return r.db.Create(pkg).Error
}

func (r *premiumRepository) GetAll() ([]entity.PremiumPackage, error) {
	var packages []entity.PremiumPackage
	err := r.db.Find(&packages).Error
	return packages, err
}
