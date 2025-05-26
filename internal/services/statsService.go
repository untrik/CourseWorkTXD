package services

import (
	"context"
	"fmt"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type statsService struct {
	statsRepo   repositories.StatsRepoInterface
	matchesRepo repositories.MatchesRepoInterface
	playersRepo repositories.PlayersRepoInterface
}

func NewStatsService(StatR repositories.StatsRepoInterface,
	MR repositories.MatchesRepoInterface, PR repositories.PlayersRepoInterface) StatsServiceInterface {
	return &statsService{
		statsRepo:   StatR,
		matchesRepo: MR,
		playersRepo: PR}
}
func (StatS *statsService) GetByPlayer(ctx context.Context, playerID int) ([]models.PlayersStatistics, error) {
	exists, err := StatS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return nil, fmt.Errorf("checking player existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("player %d not found", playerID)
	}
	return StatS.statsRepo.GetByPlayer(ctx, playerID)
}
func (StatS *statsService) AddStatistic(ctx context.Context, matchID, playerID int, kills, deaths, assists,
	creeps, denieds int16, goldPerMinute, netWorth float64) error {
	if kills < 0 || deaths < 0 || assists < 0 || creeps < 0 || denieds < 0 || goldPerMinute <= 0 || netWorth <= 0 {
		return fmt.Errorf("arguments cannot be less than 0")
	}
	existsPlayer, err := StatS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return fmt.Errorf("checking player existence: %w", err)
	}
	if !existsPlayer {
		return fmt.Errorf("player %d not found", playerID)
	}
	existsMatch, err := StatS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return fmt.Errorf("match %d not found", matchID)
	}
	return StatS.statsRepo.AddStatistic(ctx, matchID, playerID, kills, deaths, assists,
		creeps, denieds, goldPerMinute, netWorth)
}
func (StatS *statsService) DeleteByMatch(ctx context.Context, matchID int) error {
	existsMatch, err := StatS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return fmt.Errorf("match %d not found", matchID)
	}
	return StatS.statsRepo.DeleteByMatch(ctx, matchID)
}
func (StatS *statsService) DeleteByID(ctx context.Context, matchID int, playerID int) error {
	existsMatch, err := StatS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return fmt.Errorf("match %d not found", matchID)
	}
	existsPlayer, err := StatS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return fmt.Errorf("checking player existence: %w", err)
	}
	if !existsPlayer {
		return fmt.Errorf("player %d not found", playerID)
	}
	return StatS.statsRepo.DeleteByID(ctx, matchID, playerID)
}
func (StatS *statsService) GetByMatch(ctx context.Context, matchID int) ([]models.PlayersStatistics, error) {
	existsMatch, err := StatS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return nil, fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return nil, fmt.Errorf("match %d not found", matchID)
	}
	return StatS.statsRepo.GetByMatch(ctx, matchID)
}
