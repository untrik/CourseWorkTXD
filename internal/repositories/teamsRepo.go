package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type teamsRepo struct {
	DB *gorm.DB
}

func NewTeamsRepo(db *gorm.DB) *teamsRepo {
	return &teamsRepo{DB: db}
}
func (TR *teamsRepo) FindTeamByTournament(ctx context.Context, tournamentID int64) ([]models.Teams, error) {
	var teams []models.Teams
	sql := `SELECT DISTINCT tournaments.teams.* 
		FROM  tournaments.teams
		JOIN tournaments.teams_matches ON tournaments.teams.team_id = tournaments.teams_matches.team_id
		JOIN tournaments.matches ON tournaments.teams_matches.match_id = tournaments.matches.match_id
		WHERE tournaments.matches.tournament_id = ?
		ORDER BY tournaments.teams.title ASC`
	if err := TR.DB.WithContext(ctx).Raw(sql, tournamentID).Scan(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
func (TR *teamsRepo) GetAllTeams(ctx context.Context) ([]models.Teams, error) {
	var teams []models.Teams
	if err := TR.DB.WithContext(ctx).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
func (TR *teamsRepo) GetTeamsWinrates(ctx context.Context, winrate float64) ([]models.TeamsWinrate, error) {
	var teamsWinrate []models.TeamsWinrate
	sql := `
	SELECT
  	tournaments.teams.title,
  	ROUND(100 * COUNT(*) FILTER (WHERE tournaments.teams_matches.result = 'win') / COUNT(tournaments.teams_matches.match_id),2) AS winrate
	FROM tournaments.teams
	JOIN tournaments.teams_matches
 	ON tournaments.teams.team_id = tournaments.teams_matches.team_id
	WHERE tournaments.teams.team_id IN (
  	SELECT team_id
  	FROM tournaments.teams_matches)
	GROUP BY tournaments.teams.title
	HAVING
  	COUNT(*) FILTER (WHERE tournaments.teams_matches.result='win')::float
  	/ COUNT(tournaments.teams_matches.match_id) > ?
	ORDER BY winrate DESC;
`
	if err := TR.DB.WithContext(ctx).Raw(sql, winrate).Scan(&teamsWinrate).Error; err != nil {
		return nil, err
	}
	return teamsWinrate, nil
}
func (TR *teamsRepo) GetTeamsFaceToFace(ctx context.Context) ([]models.FaceToFace, error) {
	var faceToFace []models.FaceToFace
	sql := `
	SELECT 
	t1.title AS team_a,
	t2.title AS team_b,
	COUNT(*) FILTER (WHERE tm1.result = 'win')  AS wins_a,
	COUNT(*) FILTER (WHERE tm2.result = 'win')  AS wins_b,
	COUNT(*)           AS matches_played
	FROM tournaments.teams_matches AS tm1
	JOIN tournaments.teams_matches AS tm2
  	ON tm1.match_id = tm2.match_id
 	AND tm1.team_id < tm2.team_id
	JOIN tournaments.teams AS t1
  	ON tm1.team_id = t1.team_id
	JOIN tournaments.teams AS t2
  	ON tm2.team_id = t2.team_id
	GROUP BY
  	t1.title,
  	t2.title
	ORDER BY
  	t1.title,
  	t2.title;
`
	if err := TR.DB.WithContext(ctx).Raw(sql).Scan(&faceToFace).Error; err != nil {
		return nil, err
	}
	return faceToFace, nil
}
func (TR *teamsRepo) AddTeam(ctx context.Context, country string, title string, foundationDate time.Time) error {
	team := models.Teams{
		Country:        country,
		Title:          title,
		FoundationDate: foundationDate,
	}
	return TR.DB.WithContext(ctx).Create(&team).Error
}
func (TR *teamsRepo) DeleteTeam(ctx context.Context, teamID int) error {
	return TR.DB.WithContext(ctx).Delete(&models.Teams{}, teamID).Error
}
func (tr *teamsRepo) Exists(ctx context.Context, teamID int) (bool, error) {
	var count int64
	err := tr.DB.WithContext(ctx).Model(&models.Teams{}).Where("team_id = ?", teamID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (tr *teamsRepo) TeamTitleExists(ctx context.Context, title string) (bool, error) {
	var count int64
	err := tr.DB.WithContext(ctx).Model(&models.Teams{}).Where("title ILIKE ?", title).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
