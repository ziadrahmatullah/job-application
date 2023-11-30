package model

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Company  string    `gorm:"not null"`
	Quota    int       `gorm:"not null"`
	ExpireAt time.Time `gorm:"not null"`
}
