package server

import (
	patientRepository "hospital-system/internal/api/patient/repository"
	patientRoutes "hospital-system/internal/api/patient/routes"
	patientServices "hospital-system/internal/api/patient/services"
	staffRepository "hospital-system/internal/api/staff/repository"
	staffRoutes "hospital-system/internal/api/staff/routes"
	staffServices "hospital-system/internal/api/staff/services"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() *gin.Engine {
	e := gin.Default()
	staffRepository := staffRepository.NewStaffRepository(s.postgres)
	staffService := staffServices.NewStaffService(staffRepository)
	staffRoutes.RegisterStaffRoutes(e, staffService)

	patientRepository := patientRepository.NewPatientRepository(s.postgres)
	patientService := patientServices.NewPatientService(patientRepository)
	patientRoutes.RegisterPatientRoutes(e, patientService)

	return e
}
