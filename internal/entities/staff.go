package entities

import (
	"time"

	"github.com/google/uuid"
)

type Staff struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Username   string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
}
