package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type matchesRepo struct {
	DB *gorm.DB
}

func NewMatchesRepo(db *gorm.DB) *matchesRepo {
	return &matchesRepo{DB: db}
}
func (mr *matchesRepo) FindByTournament(ctx context.Context, tournamentID int64) ([]models.Matches, error) {
	var matches []models.Matches
	if err := mr.DB.WithContext(ctx).Where("tournament_id = ?", tournamentID).Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func (mr *matchesRepo) FindBeforeDate(ctx context.Context, date time.Time) ([]models.MatchesInfo, error) {

	var matches []models.MatchesInfo
	sql := `
	SELECT tournaments.teams.title, tournaments.teams_matches.result,
	tournaments.teams_matches.score ,tournaments.matches.match_date FROM tournaments.teams
	JOIN tournaments.teams_matches ON tournaments.teams.team_id = tournaments.teams_matches.team_id
	JOIN tournaments.matches ON tournaments.teams_matches.match_id = tournaments.matches.match_id
	WHERE(tournaments.matches.match_date < ?)
	ORDER BY tournaments.matches.match_date ASC,tournaments.teams.title ASC
`
	if err := mr.DB.WithContext(ctx).Raw(sql, date).Scan(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func (mr *matchesRepo) CancelMatches(ctx context.Context) error {

	sql := `
	WITH to_cancel AS (
  	SELECT mp.match_id
  	FROM tournaments.matches_participants AS mp
  	LEFT JOIN tournaments.teams_matches AS tm
    ON mp.match_id = tm.match_id
  	WHERE tm.match_id IS NULL
  	GROUP BY mp.match_id
	),
	del_parts AS (
  	DELETE FROM tournaments.matches_participants
  	WHERE match_id IN (SELECT match_id FROM to_cancel)
  	RETURNING match_id
	)
	DELETE FROM tournaments.matches
	WHERE match_id IN (SELECT match_id FROM to_cancel);
`
	if err := mr.DB.WithContext(ctx).Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
func (mr *matchesRepo) AddMatche(ctx context.Context, tournamentID int64, stageID int, matchDate time.Time) error {
	sql := `
    INSERT INTO tournaments.matches (tournament_id, stage_id, match_date)
    VALUES (?, ?, ?)
    `
	return mr.DB.WithContext(ctx).Exec(sql, tournamentID, stageID, matchDate).Error
}
func (mr *matchesRepo) FindByTeam(ctx context.Context, teamID int) ([]models.MatchesParticipants, error) {
	var matchesParticipants []models.MatchesParticipants
	if err := mr.DB.WithContext(ctx).Where("team_id = ?", teamID).Find(&matchesParticipants).Error; err != nil {
		return nil, err
	}
	return matchesParticipants, nil
}
func (mr *matchesRepo) FindByStage(ctx context.Context, stageID int) ([]models.Matches, error) {
	var matches []models.Matches
	if err := mr.DB.WithContext(ctx).Where("stage_id = ?", stageID).Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}
func (mr *matchesRepo) CountByStage(ctx context.Context) ([]models.CountByStage, error) {
	var countByStage []models.CountByStage
	sql := `
	 tournaments.tournaments_stages.stage_name,
 	COUNT(*) AS matches_count
	FROM tournaments.tournaments_stages
	JOIN tournaments.matches ON tournaments.tournaments_stages.stage_id = tournaments.matches.stage_id
	GROUP BY tournaments.tournaments_stages.stage_name
	ORDER BY matches_count DESC;
`
	if err := mr.DB.WithContext(ctx).Raw(sql).Scan(&countByStage).Error; err != nil {
		return nil, err
	}
	return countByStage, nil
}
func (mr *matchesRepo) MatchesExists(ctx context.Context, matchID int) (bool, error) {
	var count int64
	err := mr.DB.WithContext(ctx).Model(&models.Players{}).Where("match_id = ?", matchID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
