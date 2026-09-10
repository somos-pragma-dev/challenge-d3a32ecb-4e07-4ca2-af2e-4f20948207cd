package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"internal/idempotency"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idempotencyKey := idempotency.GenerateIdempotencyKey(c.Request.Context(), c.Request.Body)
	if h.checkIdempotency(idempotencyKey) {
		c.JSON(http.StatusOK, gin.H{"message": "User already exists"})
		return
	}
	h.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) checkIdempotency(key idempotency.IdempotencyKey) bool {
	// Implement idempotency check logic here
	return false
}