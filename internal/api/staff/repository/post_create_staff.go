package repository

import (
	"hospital-system/internal/entities"
)

func (r *StaffRepository) PostCreateStaff(staff entities.Staff) error {
	err := r.postgres.Create(&staff).Error
	if err != nil {
		return err
	}
	return nil
}
