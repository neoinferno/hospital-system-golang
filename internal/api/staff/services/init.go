package services

import "hospital-system/internal/api/staff"

type StaffService struct {
	repository staff.Repository
}

func NewStaffService(repository staff.Repository) *StaffService {
	return &StaffService{repository: repository}
}
