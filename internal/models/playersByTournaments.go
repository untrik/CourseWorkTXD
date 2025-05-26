package models

type PlayersByTournaments struct {
	Nickname  string `gorm:"column:nickname"`
	TeamTitle string `gorm:"column:title"`
}
