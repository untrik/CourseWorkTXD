package services

import (
	"context"
	"fmt"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type matchesService struct {
	matchesRepo     repositories.MatchesRepoInterface
	teamsRepo       repositories.TeamRepoInterface
	tournamentsRepo repositories.TournamentsRepoInterface
	stageRepo       repositories.StageRepoInterface
}

func NewMatchesService(MR repositories.MatchesRepoInterface, TR repositories.TeamRepoInterface, TourR repositories.TournamentsRepoInterface,
	SR repositories.StageRepoInterface) MatchesServiceInterface {
	return &matchesService{matchesRepo: MR,
		teamsRepo:       TR,
		tournamentsRepo: TourR,
		stageRepo:       SR}
}
func (MS *matchesService) FindByTournament(ctx context.Context, tournamentID int64) ([]models.Matches, error) {
	exists, err := MS.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("tournament %d not found", tournamentID)
	}
	return MS.matchesRepo.FindByTournament(ctx, tournamentID)
}

func (MS *matchesService) FindBeforeDate(ctx context.Context, date time.Time) ([]models.MatchesInfo, error) {
	return MS.matchesRepo.FindBeforeDate(ctx, date)
}

func (MS *matchesService) CancelMatches(ctx context.Context) error {
	return MS.matchesRepo.CancelMatches(ctx)
}

func (MS *matchesService) AddMatche(ctx context.Context, tournamentID int64, stageID int, matchDate time.Time) error {
	tournamentExists, err := MS.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("checking tournament existence: %w", err)
	}
	if !tournamentExists {
		return fmt.Errorf("tournament %d not found", tournamentID)
	}
	return MS.matchesRepo.AddMatche(ctx, tournamentID, stageID, matchDate)
}
func (MS *matchesService) FindByTeam(ctx context.Context, teamID int) ([]models.MatchesParticipants, error) {
	exists, err := MS.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("checking team existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("team %d not found", teamID)
	}
	return MS.matchesRepo.FindByTeam(ctx, teamID)

}
func (MS *matchesService) FindByStage(ctx context.Context, stageID int) ([]models.Matches, error) {
	exists, err := MS.stageRepo.StageExists(ctx, stageID)
	if err != nil {
		return nil, fmt.Errorf("checking stage existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("stage %d not found", stageID)
	}
	return MS.matchesRepo.FindByStage(ctx, stageID)
}
func (MS *matchesService) CountByStage(ctx context.Context) ([]models.CountByStage, error) {
	return MS.matchesRepo.CountByStage(ctx)
}
