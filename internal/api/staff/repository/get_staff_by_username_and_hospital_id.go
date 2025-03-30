package repository

import "hospital-system/internal/entities"

func (r *StaffRepository) GetStaffByUsernameAndHospitalId(username string, hospitalId string) (entities.Staff, error) {
	staff := entities.Staff{}
	err := r.postgres.Where("username = ? AND hospital_id = ?", username, hospitalId).First(&staff).Error
	return staff, err
}
