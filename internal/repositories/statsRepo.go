package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type statsRepo struct {
	DB *gorm.DB
}

func NewStatsRepo(db *gorm.DB) *statsRepo {
	return &statsRepo{DB: db}
}
func (StatR *statsRepo) GetByPlayer(ctx context.Context, playerID int) ([]models.PlayersStatistics, error) {
	var stats []models.PlayersStatistics
	if err := StatR.DB.WithContext(ctx).Where("player_id = ?", playerID).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}
func (StatR *statsRepo) AddStatistic(ctx context.Context, matchID, playerID int, kills, deaths, assists,
	creeps, denieds int16, goldPerMinute, netWorth float64) error {
	stats := models.PlayersStatistics{
		MatchID:       matchID,
		PlayerID:      playerID,
		Kill:          kills,
		Death:         deaths,
		Assist:        assists,
		Creep:         creeps,
		Denied:        denieds,
		GoldPerMinute: goldPerMinute,
		NetWorth:      netWorth,
	}
	return StatR.DB.WithContext(ctx).Create(&stats).Error
}
func (StatR *statsRepo) DeleteByMatch(ctx context.Context, matchID int) error {
	return StatR.DB.WithContext(ctx).Where("match_id = ?", matchID).Delete(&models.PlayersStatistics{}).Error
}
func (StatR *statsRepo) DeleteByID(ctx context.Context, matchID int, playerID int) error {
	return StatR.DB.WithContext(ctx).Delete(&models.PlayersStatistics{}, matchID, playerID).Error
}
func (StatR *statsRepo) GetByMatch(ctx context.Context, matchID int) ([]models.PlayersStatistics, error) {
	var stats []models.PlayersStatistics
	if err := StatR.DB.WithContext(ctx).Where("player_id = ?", matchID).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}
