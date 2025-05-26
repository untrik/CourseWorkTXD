package services

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
)

type StageServiceInterface interface {
	GetByTournament(ctx context.Context, tournamentID int64) ([]models.TournamentsStages, error)
	GetByID(ctx context.Context, stageID int) (*models.TournamentsStages, error)
	AddStage(ctx context.Context, tournamentID int64, stageName models.StageType) error
	Delete(ctx context.Context, stageID int) error
}
