package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type TeamRepoInterface interface {
	FindTeamByTournament(ctx context.Context, tournamentID int64) ([]models.Teams, error)
	GetAllTeams(ctx context.Context) ([]models.Teams, error)
	GetTeamsWinrates(ctx context.Context, winrate float64) ([]models.TeamsWinrate, error)
	GetTeamsFaceToFace(ctx context.Context) ([]models.FaceToFace, error)
	AddTeam(ctx context.Context, country string, title string, foundationDate time.Time) error
	DeleteTeam(ctx context.Context, teamID int) error
	Exists(ctx context.Context, teamID int) (bool, error)
	TeamTitleExists(ctx context.Context, title string) (bool, error)
}
