package repository

import "hospital-system/internal/entities"

func (r *PatientRepository) GetAllPatients(filters map[string]interface{}) ([]entities.Patient, error) {
	patients := []entities.Patient{}
	query := r.postgres
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.Find(&patients).Error
	return patients, err
}
