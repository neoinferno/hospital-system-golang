package entities

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	HospitalID   uuid.UUID `gorm:"type:uuid;not null"`
	Email        string    `gorm:"unique;not null"`
	FirstNameTh  string    `gorm:"column:first_name_th"`
	MiddleNameTh string    `gorm:"column:middle_name_th"`
	LastNameTh   string    `gorm:"column:last_name_th"`
	FirstNameEn  string    `gorm:"column:first_name_en"`
	MiddleNameEn string    `gorm:"column:middle_name_en"`
	LastNameEn   string    `gorm:"column:last_name_en"`
	DateOfBirth  string    `gorm:"column:date_of_birth"`
	PatientHN    string    `gorm:"column:patient_hn"`
	NationalID   string    `gorm:"column:national_id"`
	PassportID   string    `gorm:"column:passport_id"`
	PhoneNumber  string    `gorm:"column:phone_number"`
	CountryCode  string    `gorm:"column:country_code"`
	Gender       Gender    `gorm:"column:gender"`
	CreatedAt    time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;not null"`
}

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)
