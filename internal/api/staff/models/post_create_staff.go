package models

type PostCreateStaff struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	HospitalID string `json:"hospitalId"`
}
