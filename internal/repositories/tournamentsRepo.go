package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type tournamentsRepo struct {
	DB *gorm.DB
}

func NewTournamentsRepo(db *gorm.DB) *tournamentsRepo {
	return &tournamentsRepo{DB: db}
}
func (TourR *tournamentsRepo) TournamentExists(ctx context.Context, tournamentID int64) (bool, error) {
	var count int64
	err := TourR.DB.WithContext(ctx).Model(&models.Tournaments{}).Where("tournament_id = ?", tournamentID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (TourR *tournamentsRepo) AddTournament(ctx context.Context, tournamentName string, prizePool float64, startDate time.Time, endDate time.Time) error {
	tournament := models.Tournaments{
		TournamentName: tournamentName,
		PrizePool:      prizePool,
		StartDate:      startDate,
		EndDate:        endDate,
	}
	return TourR.DB.WithContext(ctx).Create(&tournament).Error
}
func (TourR *tournamentsRepo) GetTournamentByID(ctx context.Context, tournamentID int64) (*models.Tournaments, error) {
	var tournament models.Tournaments
	if err := TourR.DB.WithContext(ctx).Where("tournament_id = ?", tournamentID).First(&tournament).Error; err != nil {
		return nil, err
	}
	return &tournament, nil
}
func (TourR *tournamentsRepo) GetAllTournaments(ctx context.Context) ([]models.Tournaments, error) {
	var tournaments []models.Tournaments
	if err := TourR.DB.WithContext(ctx).Find(&tournaments).Error; err != nil {
		return nil, err
	}
	return tournaments, nil
}
func (TourR *tournamentsRepo) DeleteTournaments(ctx context.Context, tournamentID int64) error {
	return TourR.DB.WithContext(ctx).Delete(&models.Tournaments{}, tournamentID).Error
}
func (TourR *tournamentsRepo) TournamentNameExists(ctx context.Context, tournamentName string) (bool, error) {
	var count int64
	err := TourR.DB.WithContext(ctx).Model(&models.Tournaments{}).Where("tournament_name ILIKE ?", tournamentName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
