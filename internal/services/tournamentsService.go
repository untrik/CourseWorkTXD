package services

import (
	"context"
	"fmt"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type tournamentsService struct {
	tournamentsRepo repositories.TournamentsRepoInterface
}

func NewTournamentsService(TourR repositories.TournamentsRepoInterface) TournamentsServiceInterface {
	return &tournamentsService{tournamentsRepo: TourR}
}
func (TourR *tournamentsService) AddTournament(ctx context.Context, tournamentName string, prizePool float64, startDate time.Time, endDate time.Time) error {
	if endDate.Before(startDate) {
		return fmt.Errorf("end date (%v) must be after start date (%v)", endDate, startDate)
	}
	if prizePool < 0 {
		return fmt.Errorf("prize pool must be more then 0")
	}
	exists, err := TourR.tournamentsRepo.TournamentNameExists(ctx, tournamentName)
	if err != nil {
		return fmt.Errorf("checking tournament existence: %w", err)
	}
	if exists {
		return fmt.Errorf("tournament %s already exists", tournamentName)
	}
	return TourR.tournamentsRepo.AddTournament(ctx, tournamentName, prizePool, startDate, endDate)
}
func (TourR *tournamentsService) GetTournamentByID(ctx context.Context, tournamentID int64) (*models.Tournaments, error) {
	exists, err := TourR.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("tournament %d not found", tournamentID)
	}
	return TourR.tournamentsRepo.GetTournamentByID(ctx, tournamentID)
}
func (TourR *tournamentsService) GetAllTournaments(ctx context.Context) ([]models.Tournaments, error) {
	return TourR.tournamentsRepo.GetAllTournaments(ctx)
}
func (TourR *tournamentsService) DeleteTournaments(ctx context.Context, tournamentID int64) error {
	exists, err := TourR.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("tournament %d not found", tournamentID)
	}
	return TourR.tournamentsRepo.DeleteTournaments(ctx, tournamentID)
}
