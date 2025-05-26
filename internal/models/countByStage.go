package models

type CountByStage struct {
	StageName    string `gorm:"column:stage_name" json:"stage_name"`
	MatchesCount int    `gorm:"column:matches_count" json:"matches_count"`
}
