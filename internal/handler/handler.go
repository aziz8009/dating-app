package handler

import (
	"dating-app-backend/internal/domain/entity"
	service "dating-app-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AuthService  service.AuthService
	SwipeService service.SwipeService
}

func NewHandler(authService service.AuthService, swipeService service.SwipeService) *Handler {
	return &Handler{AuthService: authService, SwipeService: swipeService}
}

func (h *Handler) SignUp(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind the incoming JSON payload to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	if err := h.AuthService.SignUp(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, token, err := h.AuthService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
