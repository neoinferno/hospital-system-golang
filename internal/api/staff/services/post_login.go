package services

import (
	"errors"
	"hospital-system/internal/api/staff/models"
	middleware "hospital-system/internal/middlewares"

	"golang.org/x/crypto/bcrypt"
)

func (s *StaffService) PostLogin(request models.LoginRequest) (string, error) {
	staff, err := s.repository.GetStaffByUsernameAndHospitalId(request.Username, request.HospitalID)
	if err != nil {
		return "", err
	}

	if !comparePasswords(staff.Password, request.Password) {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := middleware.GenerateToken(staff.ID.String(), staff.Username, staff.HospitalID.String())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
