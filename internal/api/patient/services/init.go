package services

import (
	patient "hospital-system/internal/api/patient"
)

type PatientService struct {
	patientRepository patient.Repository
}

func NewPatientService(patientRepository patient.Repository) *PatientService {
	return &PatientService{
		patientRepository: patientRepository,
	}
}
