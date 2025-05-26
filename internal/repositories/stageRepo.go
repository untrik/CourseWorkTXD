package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type stageRepo struct {
	DB *gorm.DB
}

func NewStageRepo(db *gorm.DB) *stageRepo {
	return &stageRepo{DB: db}
}
func (SR *stageRepo) GetByTournament(ctx context.Context, tournamentID int64) ([]models.TournamentsStages, error) {
	var tournamentStages []models.TournamentsStages
	if err := SR.DB.WithContext(ctx).Where("tournament_id = ?", tournamentID).Find(&tournamentStages).Error; err != nil {
		return nil, err
	}
	return tournamentStages, nil
}
func (SR *stageRepo) GetByID(ctx context.Context, StageID int) (*models.TournamentsStages, error) {
	var tournamentStage models.TournamentsStages
	if err := SR.DB.WithContext(ctx).Where("stage_id = ?", StageID).Find(&tournamentStage).Error; err != nil {
		return nil, err
	}
	return &tournamentStage, nil
}
func (SR *stageRepo) AddStage(ctx context.Context, tournamentID int64, stageName models.StageType) error {
	tournamentStage := models.TournamentsStages{
		StageName:    stageName,
		TournamentID: tournamentID,
	}
	return SR.DB.WithContext(ctx).Create(&tournamentStage).Error
}
func (SR *stageRepo) StageExists(ctx context.Context, stageID int) (bool, error) {
	var count int64
	err := SR.DB.WithContext(ctx).Model(&models.TournamentsStages{}).Where("stage_id = ?", stageID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (SR *stageRepo) Delete(ctx context.Context, stageID int) error {
	return SR.DB.WithContext(ctx).Delete(&models.TournamentsStages{}, stageID).Error
}
