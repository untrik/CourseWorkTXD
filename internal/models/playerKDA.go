package models

import "time"

type PlayerKDA struct {
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	RoleTitle string    `gorm:"column:title" json:"title"`
	MatchDate time.Time `gorm:"column:match_date" json:"match_date"`
	KDA       float64   `gorm:"column:kda" json:"kda"`
}
