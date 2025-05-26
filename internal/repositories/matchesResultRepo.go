package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type matchesResultRepo struct {
	DB *gorm.DB
}

func NewMatchesResultRepo(db *gorm.DB) *matchesResultRepo {
	return &matchesResultRepo{DB: db}
}
func (MRR *matchesResultRepo) AddResults(ctx context.Context, matchID, teamID int, score string, result models.ResultMatch) error {
	return MRR.DB.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		scores := strings.Split(score, ":")
		revertScore := fmt.Sprintf("%s:%s", scores[1], scores[0])
		var invertResult models.ResultMatch
		switch result {
		case models.Win:
			invertResult = models.Loss
		case models.Loss:
			invertResult = models.Win
		case models.Draw:
			invertResult = models.Draw
		default:
			return fmt.Errorf("invalid result %q", result)
		}
		teamMatch1 := models.TeamsMatches{
			MatchID: matchID,
			TeamID:  teamID,
			Score:   score,
			Result:  result,
		}
		if err := db.Create(&teamMatch1).Error; err != nil {
			return err
		}
		var participants []models.MatchesParticipants
		if err := db.Where("match_id = ?", matchID).Find(&participants).Error; err != nil {
			return err
		}
		var otherTeamID int
		if participants[0].TeamID == teamID {
			otherTeamID = participants[1].TeamID
		} else if participants[1].TeamID == teamID {
			otherTeamID = participants[0].TeamID
		} else {
			return fmt.Errorf("team %d is not a participant of match %d", teamID, matchID)
		}
		teamMatch2 := models.TeamsMatches{
			MatchID: matchID,
			TeamID:  otherTeamID,
			Score:   revertScore,
			Result:  invertResult,
		}
		if err := db.Create(&teamMatch2).Error; err != nil {
			return err
		}
		return nil
	})
}
func (MRR *matchesResultRepo) GetByMatch(ctx context.Context, matchID int) ([]models.TeamsMatches, error) {
	var result []models.TeamsMatches
	if err := MRR.DB.WithContext(ctx).Where("macth_id = ?", matchID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (MRR *matchesResultRepo) GetByTeam(ctx context.Context, teamID int) ([]models.TeamsMatches, error) {
	var result []models.TeamsMatches
	if err := MRR.DB.WithContext(ctx).Where("team_id = ?", teamID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (MRR *matchesResultRepo) DeleteResults(ctx context.Context, matchID int) error {
	return MRR.DB.WithContext(ctx).Where("match_id = ?", matchID).Delete(&models.TeamsMatches{}).Error
}
