package repository

import (
	patient "hospital-system/internal/api/patient"
	"hospital-system/internal/database"

	"gorm.io/gorm"
)

type PatientRepository struct {
	postgres *gorm.DB
}

func NewPatientRepository(postgres database.Service) patient.Repository {
	repo := PatientRepository{
		postgres: postgres.DB(),
	}
	return &repo
}
