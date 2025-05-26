package models

import (
	"time"

	"gorm.io/gorm"
)

type Teams struct {
	TeamID         int            `json:"team_id" gorm:"primaryKey;autoIncrement;not null"`
	Country        string         `json:"country" gorm:"size:200;not null"`
	Title          string         `json:"title" gorm:"size:200;unique;not null"`
	FoundationDate time.Time      `json:"foundation_date" gorm:"not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
