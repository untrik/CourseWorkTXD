package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type StatsRepoInterface interface {
	GetByPlayer(ctx context.Context, playerID int) ([]models.PlayersStatistics, error)
	AddStatistic(ctx context.Context, matchID, playerID int, kills, deaths, assists,
		creeps, denieds int16, goldPerMinute, netWorth float64) error
	DeleteByMatch(ctx context.Context, matchID int) error
	DeleteByID(ctx context.Context, matchID int, playerID int) error
	GetByMatch(ctx context.Context, matchID int) ([]models.PlayersStatistics, error)
}
