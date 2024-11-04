package main

import (
	"dating-app-backend/internal/config"
	"dating-app-backend/internal/handler"
	"dating-app-backend/internal/middleware"
	"dating-app-backend/internal/repository"
	service "dating-app-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := &config.Config{
		DBUser:     "root",
		DBPassword: "",
		DBHost:     "localhost",
		DBPort:     "3306",
		DBName:     "dating_app",
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		panic("failed to connect to database")
	}

	userRepo := repository.NewUserRepository(db)
	swipeRepo := repository.NewSwipeRepository(db)     // Initialize the swipe repository
	premiumRepo := repository.NewPremiumRepository(db) // Initialize the premium repository

	authService := service.NewAuthService(userRepo)
	swipeService := service.NewSwipeService(swipeRepo, userRepo)       // Initialize swipe service
	premiumService := service.NewPremiumService(premiumRepo, userRepo) // Initialize premium service

	h := handler.NewHandler(authService, swipeService)          // Create main handler
	swipeHandler := handler.NewSwipeHandler(swipeService)       // Create swipe handler
	premiumHandler := handler.NewPremiumHandler(premiumService) // Create premium handler

	router := gin.Default()
	router.POST("/signup", h.SignUp)
	router.POST("/login", h.Login)

	// Apply auth middleware to routes that require authentication
	authorized := router.Group("/", middleware.AuthMiddleware())
	{
		authorized.POST("/swipe", swipeHandler.Swipe)
		authorized.GET("/packages", premiumHandler.GetPackages)
		authorized.POST("/purchase", premiumHandler.Purchase)
	}

	router.Run(":5000")
}
