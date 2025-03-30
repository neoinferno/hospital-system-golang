package entities

import (
	"time"

	"github.com/google/uuid"
)

type Hospital struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	NameTh    string    `gorm:"type:varchar(255);not null"`
	NameEn    string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}
