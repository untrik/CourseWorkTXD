package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type PlayersRepoInterface interface {
	FindPlayersByTournamentsPrizePool(ctx context.Context, prizePool float64) ([]models.PlayersByTournaments, error)
	UpdatePlayersRole(ctx context.Context, playerID int, roleTitle models.RolesTitle) error
	GetAllPlayers(ctx context.Context) ([]models.Players, error)
	GetPlayersByTeam(ctx context.Context, teamID int) ([]models.Players, error)
	AddPlayer(ctx context.Context, teamID int, roleID int16, nickname string, name string, lastName string, DateOfBirth time.Time) error
	DeletePlayer(ctx context.Context, playerID int) error
	FindPlayersByMatchKDA(ctx context.Context, KDA float64) ([]models.PlayerKDA, error)
	UpdatePlayerNickname(ctx context.Context, PlayerID int, nickname string) error
	PlayerExists(ctx context.Context, playerID int) (bool, error)
	NicknameExists(ctx context.Context, nickname string) (bool, error)
}
