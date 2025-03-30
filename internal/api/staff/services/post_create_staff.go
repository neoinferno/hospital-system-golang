package services

import (
	"hospital-system/internal/api/staff/models"
	"hospital-system/internal/entities"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *StaffService) PostCreateStaff(request models.PostCreateStaff) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	staff := entities.Staff{
		ID:         uuid.New(),
		Username:   request.Username,
		Password:   string(hashedPassword),
		HospitalID: uuid.MustParse(request.HospitalID),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err = h.repository.PostCreateStaff(staff)
	if err != nil {
		return err
	}
	return nil
}
