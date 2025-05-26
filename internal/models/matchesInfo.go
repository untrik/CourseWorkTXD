package models

import "time"

type MatchesInfo struct {
	TeamTitle string    `gorm:"column:title" json:"title"`
	Result    string    `gorm:"column:result" json:"result"`
	Score     string    `gorm:"column:score" json:"score"`
	MatchDate time.Time `gorm:"column:match_date" json:"match_date"`
}
