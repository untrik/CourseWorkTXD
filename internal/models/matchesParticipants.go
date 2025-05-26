package models

type MatchesParticipants struct {
	TeamID  int     `json:"team_id" gorm:"primaryKey;not null"`
	Team    Teams   `gorm:"foreignKey:TeamID"`
	MatchID int     `json:"match_id" gorm:"primaryKey;not null"`
	Match   Matches `gorm:"foreignKey:MatchID"`
}
