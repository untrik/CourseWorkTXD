package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type TournamentsRepoInterface interface {
	TournamentExists(ctx context.Context, tournamentID int64) (bool, error)
	AddTournament(ctx context.Context, tournamentName string, prizePool float64, startDate time.Time, endDate time.Time) error
	GetTournamentByID(ctx context.Context, tournamentID int64) (*models.Tournaments, error)
	GetAllTournaments(ctx context.Context) ([]models.Tournaments, error)
	DeleteTournaments(ctx context.Context, tournamentID int64) error
	TournamentNameExists(ctx context.Context, tournamentName string) (bool, error)
}
