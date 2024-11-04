package service

import (
	"dating-app-backend/internal/domain/entity"
	"dating-app-backend/internal/repository"
)

type PremiumService interface {
	PurchasePackage(userID uint, packageID uint) error
	GetPackages() ([]entity.PremiumPackage, error)
}

type premiumService struct {
	premiumRepo repository.PremiumRepository
	userRepo    repository.UserRepository
}

func NewPremiumService(premiumRepo repository.PremiumRepository, userRepo repository.UserRepository) PremiumService {
	return &premiumService{premiumRepo, userRepo}
}

func (s *premiumService) PurchasePackage(userID uint, packageID uint) error {
	// Logic to purchase the package (e.g., save to user's record)
	return nil // Implement the logic
}

func (s *premiumService) GetPackages() ([]entity.PremiumPackage, error) {
	return s.premiumRepo.GetAll()
}
