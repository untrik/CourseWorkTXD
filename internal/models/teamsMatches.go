package models

type ResultMatch string

const (
	Win  ResultMatch = "win"
	Loss ResultMatch = "loss"
	Draw ResultMatch = "draw"
)

func (rm ResultMatch) IsValid() bool {
	switch rm {
	case Win,
		Loss,
		Draw:
		return true
	default:
		return false
	}
}

type TeamsMatches struct {
	TeamID  int         `json:"team_id" gorm:"primaryKey;not null"`
	Team    Teams       `gorm:"foreignKey:TeamID"`
	MatchID int         `json:"match_id" gorm:"primaryKey;not null"`
	Match   Matches     `gorm:"foreignKey:MatchID"`
	Score   string      `json:"score" gorm:";not null;size:3"`
	Result  ResultMatch `json:"result" gorm:"size:4"`
}
