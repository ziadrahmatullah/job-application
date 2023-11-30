package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"not null"`
	CurrentJob string `gorm:"not null"`
	Age        int    `gorm:"not null"`
}
