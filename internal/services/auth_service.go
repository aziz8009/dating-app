package service

import (
	"dating-app-backend/internal/domain/entity"
	"dating-app-backend/internal/repository"
	"dating-app-backend/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user *entity.User) error
	Login(email, password string) (*entity.User, string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) SignUp(user *entity.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

// Login authenticates the user and returns a JWT token
func (s *authService) Login(email, password string) (*entity.User, string, error) {
	// Step 1: Find the user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("Data not found") // Return error if the user is not found or other DB issue
	}

	// Step 2: Check if user exists
	if user == nil {
		return nil, "", errors.New("user not found")
	}

	// Step 3: Compare the hashed password with the provided password
	errC := comparePassword(user.Password, password)
	if errC != nil {
		return nil, "", errC // Return the error from password comparison
	}

	// Step 4: Generate JWT token if password is correct
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Step 5: Return user and token
	return user, token, nil
}

func comparePassword(hashedPassword, password string) error {
	// Check if the hashed password from the database is empty
	if hashedPassword == "" {
		return errors.New("stored password hash is empty")
	}

	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		// If the comparison fails, return an "invalid password" error

		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("invalid password")
		}

		return errors.New("error comparing password hashes")
	}

	return nil
}
