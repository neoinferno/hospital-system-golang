package repository

import (
	"hospital-system/internal/api/staff"
	"hospital-system/internal/database"

	"gorm.io/gorm"
)

type StaffRepository struct {
	postgres *gorm.DB
}

func NewStaffRepository(postgres database.Service) staff.Repository {
	repo := StaffRepository{
		postgres: postgres.DB(),
	}
	return &repo
}
