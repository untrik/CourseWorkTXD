package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournaments struct {
	TournamentID   int64          `json:"tournament_id" gorm:"primaryKey;autoIncrement;not null"`
	TournamentName string         `json:"tournament_name" gorm:"size:200;unique;not null"`
	PrizePool      float64        `json:"prize_pool" gorm:"not null;check:prize_pool > 0"`
	StartDate      time.Time      `json:"start_date" gorm:"not null" `
	EndDate        time.Time      `json:"end_date" gorm:"not null;check:end_date > start_date"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
