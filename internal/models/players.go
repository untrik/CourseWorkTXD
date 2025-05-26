package models

import (
	"time"

	"gorm.io/gorm"
)

type Players struct {
	PlayerID    int            `json:"player_id" gorm:"primaryKey;autoIncrement;not null"`
	TeamID      int            `json:"team_id" gorm:"not null"`
	Team        Teams          `gorm:"foreignKey:TeamID"`
	RoleID      int16          `json:"role_id" gorm:"not null"`
	Role        Roles          `gorm:"foreignKey:RoleID"`
	Nickname    string         `json:"nickname" gorm:"size:200;unique;not null"`
	Name        string         `json:"name" gorm:"size:200;not null"`
	LastName    string         `json:"lastname" gorm:"size:200;not null"`
	DateOfBirth time.Time      `json:"date_of_birth" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
