package services

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type MatchesResultServiceInterface interface {
	AddResults(ctx context.Context, matchID, teamID int, score string, result models.ResultMatch) error
	GetByMatch(ctx context.Context, matchID int) ([]models.TeamsMatches, error)
	GetByTeam(ctx context.Context, teamID int) ([]models.TeamsMatches, error)
	DeleteResults(ctx context.Context, matchID int) error
}
