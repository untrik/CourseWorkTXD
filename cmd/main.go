package main

import (
	"log"

	"github.com/untrik/CourseWorkTXD/cmd/config"
	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(config.GetDatabaseURL(cfg)))
	if err != nil {
		return
	}
	repos := repositories.NewRepoFactory(db)
	err = db.AutoMigrate(
		&models.Matches{},
		&models.MatchesParticipants{},
		&models.Players{},
		&models.PlayersStatistics{},
		&models.Roles{},
		&models.Teams{},
		&models.TeamsMatches{},
		&models.Tournaments{},
		&models.TournamentsStages{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции таблиц:", err)
	}

}
