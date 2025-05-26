package services

import (
	"context"
	"fmt"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type participantsService struct {
	participantsRepo repositories.ParticipantsRepoInterface
	teamsRepo        repositories.TeamRepoInterface
	matchesRepo      repositories.MatchesRepoInterface
	tournamentsRepo  repositories.TournamentsRepoInterface
}

func NewParticipantsService(PaR repositories.ParticipantsRepoInterface, TR repositories.TeamRepoInterface,
	MR repositories.MatchesRepoInterface, TouR repositories.TournamentsRepoInterface) ParticipantsServiceInterface {
	return &participantsService{
		participantsRepo: PaR,
		matchesRepo:      MR,
		teamsRepo:        TR,
		tournamentsRepo:  TouR}
}

func (PaS *participantsService) GetByTournament(ctx context.Context, tournamentID int64) ([]models.ParticipantsInfo, error) {
	exists, err := PaS.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("tournament %d not found", tournamentID)
	}
	return PaS.participantsRepo.GetByTournament(ctx, tournamentID)
}
func (PaS *participantsService) GetByMatch(ctx context.Context, matchID int) ([]models.MatchesParticipants, error) {
	exists, err := PaS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return nil, fmt.Errorf("checking matches existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("matches %d not found", matchID)
	}
	return PaS.participantsRepo.GetByMatch(ctx, matchID)

}
func (PaS *participantsService) AddParticipant(ctx context.Context, matchID, teamID int) error {
	existsMatches, err := PaS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatches {
		return fmt.Errorf("match %d not found", matchID)
	}
	existsTeam, err := PaS.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return fmt.Errorf("team matches existence: %w", err)
	}
	if !existsTeam {
		return fmt.Errorf("team %d not found", teamID)
	}
	return PaS.AddParticipant(ctx, matchID, teamID)
}
func (PaS *participantsService) DeleteParticipant(ctx context.Context, matchID, teamID int) error {
	existsMatches, err := PaS.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatches {
		return fmt.Errorf("match %d not found", matchID)
	}
	existsTeam, err := PaS.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return fmt.Errorf("team matches existence: %w", err)
	}
	if !existsTeam {
		return fmt.Errorf("team %d not found", teamID)
	}
	return PaS.DeleteParticipant(ctx, matchID, teamID)
}
