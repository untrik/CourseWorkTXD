package services

import (
	"context"
	"fmt"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/repositories"
)

type matchesResultService struct {
	matchesResultRepo repositories.MatchesResultRepoInterface
	matchesRepo       repositories.MatchesRepoInterface
	teamsRepo         repositories.TeamRepoInterface
}

func NewMatchesResultService(MRR repositories.MatchesResultRepoInterface, MR repositories.MatchesRepoInterface,
	TR repositories.TeamRepoInterface) MatchesResultServiceInterface {
	return &matchesResultService{
		matchesResultRepo: MRR,
		matchesRepo:       MR,
		teamsRepo:         TR,
	}
}
func (MRR *matchesResultService) AddResults(ctx context.Context, matchID, teamID int, score string, result models.ResultMatch) error {
	if !result.IsValid() {
		return fmt.Errorf("unknown result name")
	}
	if len(score) != 3 {
		return fmt.Errorf("the score must be 3 in length")
	}
	existsMatch, err := MRR.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return fmt.Errorf("match %d not found", matchID)
	}
	existsTeam, err := MRR.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return fmt.Errorf("checking team existence: %w", err)
	}
	if !existsTeam {
		return fmt.Errorf("team %d not found", teamID)
	}

	return MRR.matchesResultRepo.AddResults(ctx, matchID, teamID, score, result)
}
func (MRR *matchesResultService) GetByMatch(ctx context.Context, matchID int) ([]models.TeamsMatches, error) {
	existsMatch, err := MRR.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return nil, fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return nil, fmt.Errorf("match %d not found", matchID)
	}
	return MRR.matchesResultRepo.GetByMatch(ctx, matchID)
}
func (MRR *matchesResultService) GetByTeam(ctx context.Context, teamID int) ([]models.TeamsMatches, error) {
	existsTeam, err := MRR.teamsRepo.Exists(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("checking team existence: %w", err)
	}
	if !existsTeam {
		return nil, fmt.Errorf("team %d not found", teamID)
	}
	return MRR.matchesResultRepo.GetByTeam(ctx, teamID)
}
func (MRR *matchesResultService) DeleteResults(ctx context.Context, matchID int) error {
	existsMatch, err := MRR.matchesRepo.MatchesExists(ctx, matchID)
	if err != nil {
		return fmt.Errorf("checking match existence: %w", err)
	}
	if !existsMatch {
		return fmt.Errorf("match %d not found", matchID)
	}
	return MRR.matchesResultRepo.DeleteResults(ctx, matchID)
}
