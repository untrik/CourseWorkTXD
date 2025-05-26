package models

type PlayersStatistics struct {
	PlayerID      int     `json:"player_id" gorm:"primaryKey;not null"`
	Player        Players `gorm:"foreignKey:PlayerID"`
	MatchID       int     `json:"match_id" gorm:"primaryKey;not null"`
	Match         Matches `gorm:"foreignKey:MatchID"`
	Kill          int16   `json:"kill" gorm:"default:0"`
	Death         int16   `json:"death" gorm:"default:0"`
	Assist        int16   `json:"assist" gorm:"default:0"`
	GoldPerMinute float64 `json:"gold_per_minute" gorm:"default:0"`
	NetWorth      float64 `json:"net_worth" gorm:"default:0"`
	Creep         int16   `json:"creep" gorm:"default:0"`
	Denied        int16   `json:"denied" gorm:"default:0"`
}
