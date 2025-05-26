package repositories

import (
	"context"
	"time"

	"github.com/untrik/CourseWorkTXD/internal/models"
	"gorm.io/gorm"
)

type playersRepo struct {
	DB *gorm.DB
}

func NewPlayersRepo(db *gorm.DB) *playersRepo {
	return &playersRepo{DB: db}
}
func (PR *playersRepo) FindPlayersByTournamentsPrizePool(ctx context.Context, prizePool float64) ([]models.PlayersByTournaments, error) {
	var players []models.PlayersByTournaments
	sql := `
	SELECT DISTINCT tournaments.players.nickname, tournaments.teams.title  FROM tournaments.players
	JOIN tournaments.teams ON tournaments.players.team_id = tournaments.teams.team_id
	JOIN tournaments.matches_participants ON tournaments.teams.team_id = tournaments.matches_participants.team_id
	JOIN tournaments.matches ON tournaments.matches_participants.match_id = tournaments.matches.match_id
	JOIN tournaments.tournaments ON tournaments.matches.tournament_id = tournaments.tournaments.tournament_id
	WHERE(tournaments.tournaments.prize_pool >= ? )
	ORDER BY players.nickname ASC, tournaments.teams.title ASC
`
	if err := PR.DB.WithContext(ctx).Raw(sql, prizePool).Scan(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
func (PR *playersRepo) UpdatePlayersRole(ctx context.Context, playerID int, roleTitle models.RolesTitle) error {
	sql := `UPDATE tournaments.players
	SET role_id = (
  	SELECT id
  	FROM tournaments.roles
  	WHERE title = ?)
	WHERE player_id = ?;
`
	return PR.DB.WithContext(ctx).Exec(sql, roleTitle, playerID).Error
}
func (PR *playersRepo) GetAllPlayers(ctx context.Context) ([]models.Players, error) {
	var players []models.Players
	if err := PR.DB.WithContext(ctx).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
func (PR *playersRepo) GetPlayersByTeam(ctx context.Context, teamID int) ([]models.Players, error) {
	var players []models.Players
	if err := PR.DB.WithContext(ctx).Where("team_id = ?", teamID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
func (PR *playersRepo) AddPlayer(ctx context.Context, teamID int, roleID int16, nickname string, name string, lastName string, DateOfBirth time.Time) error {
	player := models.Players{
		TeamID:   teamID,
		RoleID:   roleID,
		Nickname: nickname,
		Name:     name,
		LastName: lastName,
	}
	return PR.DB.Create(&player).Error
}

func (PR *playersRepo) DeletePlayer(ctx context.Context, playerID int) error {
	return PR.DB.WithContext(ctx).Delete(&models.Players{}, playerID).Error
}
func (PR *playersRepo) FindPlayersByMatchKDA(ctx context.Context, KDA float64) ([]models.PlayerKDA, error) {
	var players []models.PlayerKDA
	sql := `
	SELECT *
	FROM(
	SELECT tournaments.players.nickname, tournaments.roles.title,
	tournaments.matches.match_date,
	ROUND(
	CASE
	WHEN tournaments.players_statistics.death = 0 
		THEN tournaments.players_statistics.kill + tournaments.players_statistics.assist
	ELSE (tournaments.players_statistics.kill + tournaments.players_statistics.assist)::numeric / tournaments.players_statistics.death
	END ,1)AS kda
	FROM tournaments.roles
	JOIN tournaments.players ON tournaments.roles.id = tournaments.players.role_id
	JOIN tournaments.players_statistics ON tournaments.players.player_id = tournaments.players_statistics.player_id
	JOIN tournaments.matches ON tournaments.matches.match_id = tournaments.players_statistics.match_id)
	WHERE (kda >= 2)
	ORDER BY match_date ASC,kda ASC
`
	if err := PR.DB.WithContext(ctx).Raw(sql, KDA).Scan(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
func (PR *playersRepo) UpdatePlayerNickname(ctx context.Context, playerID int, nickname string) error {
	var player models.Players
	if err := PR.DB.WithContext(ctx).First(&player, playerID).Error; err != nil {
		return err
	}
	if err := PR.DB.WithContext(ctx).Where("player_id = ?", playerID).Update("nickname", nickname).Error; err != nil {
		return err
	}
	return nil
}
func (PR *playersRepo) PlayerExists(ctx context.Context, playerID int) (bool, error) {
	var count int64
	err := PR.DB.WithContext(ctx).Model(&models.Players{}).Where("player_id = ?", playerID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (PR *playersRepo) NicknameExists(ctx context.Context, nickname string) (bool, error) {
	var count int64
	err := PR.DB.WithContext(ctx).Model(&models.Players{}).Where("nickname ILIKE ?", nickname).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
