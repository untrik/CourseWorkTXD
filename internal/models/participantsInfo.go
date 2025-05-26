package models

import "time"

type ParticipantsInfo struct {
	TeamID    int       `gorm:"column:team_id" json:"team_id"`
	MatchID   int       `gorm:"column:match_id" json:"match_id"`
	StageID   int       `gorm:"column:stage_id" json:"stage_id"`
	MatchDate time.Time `gorm:"column:match_date" json:"match_date"`
}
