package models

type TeamsWinrate struct {
	Title   string  `gorm:"column:title" json:"title"`
	Winrate float64 `gorm:"column:winrate" json:"winrate"`
}
