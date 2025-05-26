package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type MatchesRepoInterface interface {
	FindByTeam(ctx context.Context, teamID int) ([]models.MatchesParticipants, error)
	FindByTournament(ctx context.Context, tournamentID int64) ([]models.Matches, error)
	FindBeforeDate(ctx context.Context, date time.Time) ([]models.MatchesInfo, error)
	CancelMatches(ctx context.Context) error
	AddMatche(ctx context.Context, tournamentID int64, stageID int, matchDate time.Time) error
	FindByStage(ctx context.Context, stageID int) ([]models.Matches, error)
	CountByStage(ctx context.Context) ([]models.CountByStage, error)
	MatchesExists(ctx context.Context, matchID int) (bool, error)
}
