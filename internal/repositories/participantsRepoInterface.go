package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type ParticipantsRepoInterface interface {
	GetByTournament(ctx context.Context, tournamentID int64) ([]models.ParticipantsInfo, error)
	GetByMatch(ctx context.Context, matchID int) ([]models.MatchesParticipants, error)
	AddParticipant(ctx context.Context, matchID, teamID int) error
	DeleteParticipant(ctx context.Context, matchID, teamID int) error
}
