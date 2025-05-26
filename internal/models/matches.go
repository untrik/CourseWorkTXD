package models

import (
	"time"
)

type Matches struct {
	MatchID      int               `json:"match_id" gorm:"primaryKey;autoIncrement;not null"`
	TournamentID int64             `json:"tournament_id" gorm:"not null"`
	Tournament   Tournaments       `gorm:"foreignKey:TournamentID"`
	StageID      int               `json:"stage_id" gorm:"not null"`
	Stage        TournamentsStages `gorm:"foreignKey:StageID"`
	MatchDate    time.Time         `json:"match_date" gorm:"not null"`
}
