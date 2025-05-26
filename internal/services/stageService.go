package services

import (
	"context"
	"fmt"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type stageService struct {
	stageRepo       repositories.StageRepoInterface
	tournamentsRepo repositories.TournamentsRepoInterface
}

func NewStageService(SR repositories.StageRepoInterface, TR repositories.TournamentsRepoInterface) StageServiceInterface {
	return &stageService{stageRepo: SR, tournamentsRepo: TR}
}
func (SR *stageService) GetByTournament(ctx context.Context, tournamentID int64) ([]models.TournamentsStages, error) {
	exists, err := SR.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("tournament %d not found", tournamentID)
	}
	return SR.stageRepo.GetByTournament(ctx, tournamentID)
}
func (SR *stageService) GetByID(ctx context.Context, stageID int) (*models.TournamentsStages, error) {
	exists, err := SR.stageRepo.StageExists(ctx, stageID)
	if err != nil {
		return nil, fmt.Errorf("checking stage existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("stage %d not found", stageID)
	}
	return SR.stageRepo.GetByID(ctx, stageID)
}
func (SR *stageService) AddStage(ctx context.Context, tournamentID int64, stageName models.StageType) error {
	if !stageName.IsValid() {
		return fmt.Errorf("unknown stage name")
	}
	existsTournament, err := SR.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("checking tournament existence: %w", err)
	}
	if !existsTournament {
		return fmt.Errorf("tournament %d not found", tournamentID)
	}
	return SR.stageRepo.AddStage(ctx, tournamentID, stageName)
}
func (SR *stageService) Delete(ctx context.Context, stageID int) error {
	exists, err := SR.stageRepo.StageExists(ctx, stageID)
	if err != nil {
		return fmt.Errorf("checking stage existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("stage %d not found", stageID)
	}
	return SR.stageRepo.Delete(ctx, stageID)
}
