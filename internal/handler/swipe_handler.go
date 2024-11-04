package handler

import (
	service "dating-app-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SwipeHandler struct {
	SwipeService service.SwipeService
}

func NewSwipeHandler(swipeService service.SwipeService) *SwipeHandler {
	return &SwipeHandler{SwipeService: swipeService}
}

// Handler for swiping
func (h *SwipeHandler) Swipe(c *gin.Context) {
	var input struct {
		TargetUserID uint   `json:"target_user_id"`
		Action       string `json:"action"` // "like" or "pass"
	}

	userID, exists := c.Get("userID") // Assuming you set userID in the context after login
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.SwipeService.Swipe(userID.(uint), input.TargetUserID, input.Action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "swipe recorded successfully"})
}
