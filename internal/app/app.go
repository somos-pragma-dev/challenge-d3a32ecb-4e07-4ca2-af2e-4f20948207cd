package app

import (
	"github.com/gin-gonic/gin"
	"internal/domain"
	"internal/infrastructure"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	userHandler := &domain.UserHandler{DB: db}
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}