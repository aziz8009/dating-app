package handler

import (
	service "dating-app-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PremiumHandler struct {
	PremiumService service.PremiumService
}

func NewPremiumHandler(premiumService service.PremiumService) *PremiumHandler {
	return &PremiumHandler{PremiumService: premiumService}
}

// Handler to get all premium packages
func (h *PremiumHandler) GetPackages(c *gin.Context) {
	packages, err := h.PremiumService.GetPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// Handler to purchase a package
func (h *PremiumHandler) Purchase(c *gin.Context) {
	var input struct {
		PackageID uint `json:"package_id"`
	}

	userID, exists := c.Get("userID") // Assuming you've set the user ID in context after login
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PremiumService.PurchasePackage(userID.(uint), input.PackageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "package purchased successfully"})
}
