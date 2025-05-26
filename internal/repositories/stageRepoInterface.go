package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type StageRepoInterface interface {
	GetByTournament(ctx context.Context, tournamentID int64) ([]models.TournamentsStages, error)
	GetByID(ctx context.Context, StageID int) (*models.TournamentsStages, error)
	AddStage(ctx context.Context, tournamentID int64, stageName models.StageType) error
	StageExists(ctx context.Context, stageID int) (bool, error)
	Delete(ctx context.Context, stageID int) error
}
