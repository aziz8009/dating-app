package service

import (
	"dating-app-backend/internal/domain/entity"
	"dating-app-backend/internal/repository"
	"errors"
	"time"
)

type SwipeService interface {
	Swipe(userID uint, profileID uint, action string) error
}

type swipeService struct {
	swipeRepo repository.SwipeRepository
	userRepo  repository.UserRepository
}

func NewSwipeService(swipeRepo repository.SwipeRepository, userRepo repository.UserRepository) SwipeService {
	return &swipeService{swipeRepo, userRepo}
}

func (s *swipeService) Swipe(userID uint, profileID uint, action string) error {
	// Check the number of swipes made today by the user
	swipesToday, err := s.swipeRepo.CountTodaySwipes(userID)
	if err != nil {
		return err
	}

	if swipesToday >= 10 {
		return errors.New("swipe limit reached for today")
	}

	// Create the swipe record
	swipe := &entity.Swipe{
		UserID:    userID,
		ProfileID: profileID,
		Action:    action,
		CreatedAt: time.Now(),
	}

	return s.swipeRepo.Create(swipe)
}
