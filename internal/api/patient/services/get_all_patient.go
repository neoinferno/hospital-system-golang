package services

import "hospital-system/internal/entities"

func (s *PatientService) GetAllPatients(filters map[string]interface{}) ([]entities.Patient, error) {
	patients, err := s.patientRepository.GetAllPatients(filters)
	if err != nil {
		return nil, err
	}
	return patients, nil
}
