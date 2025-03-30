package hospital

import "hospital-system/internal/entities"

type Service interface {
	GetPatientById(id string) (string, error)
	GetAllPatients(filters map[string]interface{}) ([]entities.Patient, error)
}
type Repository interface {
	GetAllPatients(filters map[string]interface{}) ([]entities.Patient, error)
}
