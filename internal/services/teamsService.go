package services

import (
	"context"
	"fmt"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type teamsService struct {
	teamsRepo       repositories.TeamRepoInterface
	tournamentsRepo repositories.TournamentsRepoInterface
}

func NewTeamsService(TR repositories.TeamRepoInterface, TourR repositories.TournamentsRepoInterface) TeamsServiceInterface {
	return &teamsService{teamsRepo: TR,
		tournamentsRepo: TourR}
}
func (ts *teamsService) FindTeamByTournament(ctx context.Context, tournamentID int64) ([]models.Teams, error) {
	exists, err := ts.tournamentsRepo.TournamentExists(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("checking tournament existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("tournament %d not found", tournamentID)
	}
	return ts.teamsRepo.FindTeamByTournament(ctx, tournamentID)
}
func (ts *teamsService) GetAllTeams(ctx context.Context) ([]models.Teams, error) {
	return ts.teamsRepo.GetAllTeams(ctx)
}
func (ts *teamsService) GetTeamsWinrates(ctx context.Context, winrate float64) ([]models.TeamsWinrate, error) {
	if winrate < 0 || winrate > 100 {
		return nil, fmt.Errorf("winrate %.2f must be between 0 and 100", winrate)
	}
	winrate = winrate / 100
	return ts.teamsRepo.GetTeamsWinrates(ctx, winrate)
}
func (ts *teamsService) GetTeamsFaceToFace(ctx context.Context) ([]models.FaceToFace, error) {
	return ts.teamsRepo.GetTeamsFaceToFace(ctx)
}
func (ts *teamsService) AddTeam(ctx context.Context, country string, title string, foundationDate time.Time) error {
	existsTeam, err := ts.teamsRepo.TeamTitleExists(ctx, title)
	if err != nil {
		return fmt.Errorf("checking team title existence: %w", err)
	}
	if existsTeam {
		return fmt.Errorf("team title %s already exists", title)
	}
	return ts.teamsRepo.AddTeam(ctx, country, title, foundationDate)
}
func (ts *teamsService) DeleteTeam(ctx context.Context, teamID int) error {
	exists, err := ts.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return fmt.Errorf("checking team existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("team %d not found", teamID)
	}
	return ts.teamsRepo.DeleteTeam(ctx, teamID)
}
