package model

import (
	"time"

	"gorm.io/gorm"
)

type JobApply struct {
	gorm.Model
	UserId    uint      `gorm:"not null" json:"user_id"`
	JobId     uint      `gorm:"not null" json:"job_id"`
	AppliedAt time.Time `gorm:"not null"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
	Job       Job       `gorm:"foreignKey:job_id;references:id"`
}
