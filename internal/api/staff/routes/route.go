package routes

import (
	hospital "hospital-system/internal/api/staff"
	"hospital-system/internal/api/staff/handlers"
	middleware "hospital-system/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterStaffRoutes(r *gin.Engine, service hospital.Service) *gin.Engine {
	handler := handlers.NewHandlers(service)
	r.POST("/api/staffs/login", handler.PostLogin)

	staffApi := r.Group("/api/staffs")
	staffApi.Use(middleware.JWTAuthMiddleware())
	{
		staffApi.POST("", handler.PostCreateStaff)
	}

	return r
}
