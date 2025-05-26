package services

import (
	"context"
	"fmt"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type playersService struct {
	playersRepo repositories.PlayersRepoInterface
	teamsRepo   repositories.TeamRepoInterface
}

func NewPlayersService(PR repositories.PlayersRepoInterface, TR repositories.TeamRepoInterface) PlayersServiceInterface {
	return &playersService{playersRepo: PR,
		teamsRepo: TR}
}
func (PS *playersService) FindPlayersByTournamentsPrizePool(ctx context.Context, prizePool float64) ([]models.PlayersByTournaments, error) {
	if prizePool <= 0 {
		return nil, fmt.Errorf("the prize pool cannot be less than or equal to 0")
	}
	return PS.playersRepo.FindPlayersByTournamentsPrizePool(ctx, prizePool)
}
func (PS *playersService) UpdatePlayersRole(ctx context.Context, playerID int, roleTitle models.RolesTitle) error {
	if !roleTitle.IsValid() {
		return fmt.Errorf("unknown role name")
	}
	exists, err := PS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return fmt.Errorf("checking player existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("player %d not found", playerID)
	}

	return PS.playersRepo.UpdatePlayersRole(ctx, playerID, roleTitle)
}
func (PS *playersService) GetAllPlayers(ctx context.Context) ([]models.Players, error) {
	return PS.playersRepo.GetAllPlayers(ctx)
}
func (PS *playersService) GetPlayersByTeam(ctx context.Context, teamID int) ([]models.Players, error) {
	exists, err := PS.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("checking team existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("team %d not found", teamID)
	}
	return PS.playersRepo.GetPlayersByTeam(ctx, teamID)
}
func (PS *playersService) AddPlayer(ctx context.Context, teamID int, roleID int16, nickname string, name string, lastName string, DateOfBirth time.Time) error {
	existsTeam, err := PS.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return fmt.Errorf("checking team existence: %w", err)
	}
	if !existsTeam {
		return fmt.Errorf("team %d not found", teamID)
	}
	existsNickname, err := PS.playersRepo.NicknameExists(ctx, nickname)
	if err != nil {
		return fmt.Errorf("checking nickname existence: %w", err)
	}
	if existsNickname {
		return fmt.Errorf("nickname %s is already exists", nickname)
	}
	return PS.playersRepo.AddPlayer(ctx, teamID, roleID, nickname, name, lastName, DateOfBirth)
}
func (PS *playersService) DeletePlayer(ctx context.Context, playerID int) error {
	exists, err := PS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return fmt.Errorf("checking player existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("player %d not found", playerID)
	}
	return PS.playersRepo.DeletePlayer(ctx, playerID)
}
func (PS *playersService) FindPlayersByMatchKDA(ctx context.Context, KDA float64) ([]models.PlayerKDA, error) {
	if KDA < 0 {
		return nil, fmt.Errorf("the KDA cannot be less than 0")
	}
	return PS.playersRepo.FindPlayersByMatchKDA(ctx, KDA)
}
func (PS *playersService) UpdatePlayerNickname(ctx context.Context, playerID int, nickname string) error {
	existsPlayer, err := PS.playersRepo.PlayerExists(ctx, playerID)
	if err != nil {
		return fmt.Errorf("checking player existence: %w", err)
	}
	if !existsPlayer {
		return fmt.Errorf("player %d not found", playerID)
	}
	existsNickname, err := PS.playersRepo.NicknameExists(ctx, nickname)
	if err != nil {
		return fmt.Errorf("checking nickname existence: %w", err)
	}
	if existsNickname {
		return fmt.Errorf("nickname %s is already exists", nickname)
	}

	return PS.playersRepo.UpdatePlayerNickname(ctx, playerID, nickname)
}
