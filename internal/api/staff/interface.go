package staff

import (
	"hospital-system/internal/api/staff/models"
	"hospital-system/internal/entities"
)

type Service interface {
	PostCreateStaff(request models.PostCreateStaff) error
	PostLogin(request models.LoginRequest) (string, error)
}
type Repository interface {
	PostCreateStaff(staff entities.Staff) error
	GetStaffByUsernameAndHospitalId(username string, hospitalId string) (entities.Staff, error)
}
