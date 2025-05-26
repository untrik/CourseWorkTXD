package models

type FaceToFace struct {
	TeamA         string `gorm:"column:team_a" json:"team_a"`
	TeamB         string `gorm:"column:team_b" json:"team_b"`
	WinsA         int    `gorm:"column:wins_a" json:"wins_a"`
	WinsB         int    `gorm:"column:wins_b" json:"wins_b"`
	MatchesPlayed int    `gorm:"column:matches_played" json:"matches_played"`
}
