package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Company string `gorm:"not null"`
	Quota   int    `gorm:"not null"`
}
