package routes

import (
	hospital "hospital-system/internal/api/patient"
	"hospital-system/internal/api/patient/handlers"
	middleware "hospital-system/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(r *gin.Engine, service hospital.Service) *gin.Engine {
	handler := handlers.NewHandlers(service)
	patientApi := r.Group("/api/patients")

	patientApi.Use(middleware.JWTAuthMiddleware())
	{
		patientApi.GET("", handler.GetAllPatients)
	}

	return r
}
