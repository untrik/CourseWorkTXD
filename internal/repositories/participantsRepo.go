package repositories

import (
	"context"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type participantsRepo struct {
	DB *gorm.DB
}

func NewParticipantsRepo(db *gorm.DB) *participantsRepo {
	return &participantsRepo{DB: db}
}
func (PaR *participantsRepo) GetByTournament(ctx context.Context, tournamentID int64) ([]models.ParticipantsInfo, error) {
	var participants []models.ParticipantsInfo
	sql := `
	SELECT team_id, tournaments.matches_participants.match_id, stage_id, match_date FROM tournaments.matches_participants
	JOIN tournaments.matches ON tournaments.matches_participants.match_id = tournaments.matches.match_id
	JOIN tournaments.tournaments ON tournaments.matches.tournament_id = tournaments.tournaments.tournament_id
	WHERE tournaments.tournaments.tournament_id = ?
	ORDER BY match_date ASC`
	if err := PaR.DB.WithContext(ctx).Raw(sql, tournamentID).Scan(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}
func (PaR *participantsRepo) GetByMatch(ctx context.Context, matchID int) ([]models.MatchesParticipants, error) {
	var participants []models.MatchesParticipants
	if err := PaR.DB.WithContext(ctx).Where("match_id = ?", matchID).Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}
func (PaR *participantsRepo) AddParticipant(ctx context.Context, matchID, teamID int) error {
	participants := models.MatchesParticipants{
		MatchID: matchID,
		TeamID:  teamID,
	}
	return PaR.DB.WithContext(ctx).Create(participants).Error
}
func (PaR *participantsRepo) DeleteParticipant(ctx context.Context, matchID, teamID int) error {
	return PaR.DB.WithContext(ctx).Delete(&models.MatchesParticipants{}, matchID, teamID).Error
}
