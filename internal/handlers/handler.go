package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/untrik/CourseWorkTXD/internal/models"
	"github.com/untrik/CourseWorkTXD/internal/services"
)

// Matches
func GetMatchesByTeam(svc services.MatchesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		matches, err := svc.FindByTeam(c.Request().Context(), teamID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}
func GetMatchesByTournament(svc services.MatchesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		matches, err := svc.FindByTournament(c.Request().Context(), tournamentID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}
func GetMatchesByStage(svc services.MatchesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		stageID, err := strconv.Atoi(c.Param("stageID"))
		if err != nil || stageID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid stageID"})
		}
		matches, err := svc.FindByStage(c.Request().Context(), stageID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}
func GetMatchesBeforeDate(svc services.MatchesServiceInterface) echo.HandlerFunc {
	type request struct {
		Date time.Time `json:"date" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		matches, err := svc.FindBeforeDate(c.Request().Context(), req.Date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}

func DeleteMatches(svc services.MatchesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := svc.CancelMatches(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetCountByStage(svc services.MatchesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		countByStage, err := svc.CountByStage(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, countByStage)
	}
}
func AddMatche(svc services.MatchesServiceInterface) echo.HandlerFunc {
	type request struct {
		StageID   int       `json:"stage_id" validate:"required gt=0"`
		MatchDate time.Time `json:"match_date" validate:"required"`
	}
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		if err := svc.AddMatche(c.Request().Context(), tournamentID, req.StageID, req.MatchDate); err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}

// Tournaments
func AddTournament(svc services.TournamentsServiceInterface) echo.HandlerFunc {
	type request struct {
		TournamentName string    `json:"tournament_name" validate:"required"`
		PrizePool      float64   `json:"prize_pool" validate:"required gt=0"`
		StartDate      time.Time `json:"start_date" validate:"required"`
		EndDate        time.Time `json:"end_date" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddTournament(c.Request().Context(), req.TournamentName, req.PrizePool,
			req.StartDate, req.EndDate); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func GetTournamentByID(svc services.TournamentsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		tournaments, err := svc.GetTournamentByID(c.Request().Context(), tournamentID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, tournaments)
	}
}
func GetAllTournaments(svc services.TournamentsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournaments, err := svc.GetAllTournaments(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, tournaments)
	}
}
func DeleteTournament(svc services.TournamentsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		if err := svc.DeleteTournaments(c.Request().Context(), tournamentID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}

// MatchesResult

func AddResults(svc services.MatchesResultServiceInterface) echo.HandlerFunc {
	type request struct {
		Score  string             `json:"score" validate:"required"`
		Result models.ResultMatch `json:"result" validate:"required"`
	}
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddResults(c.Request().Context(), matchID, teamID, req.Score, req.Result); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func GetMatchesResultByMatch(svc services.MatchesResultServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		matches, err := svc.GetByMatch(c.Request().Context(), matchID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}
func GetMatchesResultByTeam(svc services.MatchesResultServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		matches, err := svc.GetByTeam(c.Request().Context(), teamID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, matches)
	}
}
func DeleteResults(svc services.MatchesResultServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		if err := svc.DeleteResults(c.Request().Context(), matchID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}

// Participants
func GetByTournament(svc services.ParticipantsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		participants, err := svc.GetByTournament(c.Request().Context(), tournamentID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		return c.JSON(http.StatusOK, participants)
	}
}
func GetByMatch(svc services.ParticipantsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		participants, err := svc.GetByMatch(c.Request().Context(), matchID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, participants)
	}
}
func AddParticipant(svc services.ParticipantsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		if err := svc.AddParticipant(c.Request().Context(), matchID, teamID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteParticipant(svc services.ParticipantsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		if err := svc.DeleteParticipant(c.Request().Context(), matchID, teamID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}

// Players
func FindPlayersByTournamentsPrizePool(svc services.PlayersServiceInterface) echo.HandlerFunc {
	type request struct {
		PrizePool float64 `json:"prize_pool" validate:"required gt=0"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		players, err := svc.FindPlayersByTournamentsPrizePool(c.Request().Context(), req.PrizePool)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, players)
	}
}
func UpdatePlayersRole(svc services.PlayersServiceInterface) echo.HandlerFunc {
	type request struct {
		RoleTitle models.RolesTitle `json:"role_title" validate:"required"`
	}
	return func(c echo.Context) error {
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.UpdatePlayersRole(c.Request().Context(), playerID, req.RoleTitle); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetAllPlayers(svc services.PlayersServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		players, err := svc.GetAllPlayers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, players)
	}
}
func GetPlayersByTeam(svc services.PlayersServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		players, err := svc.GetPlayersByTeam(c.Request().Context(), teamID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, players)
	}
}
func AddPlayer(svc services.PlayersServiceInterface) echo.HandlerFunc {
	type request struct {
		TeamID      int       `json:"team_id" validate:"required gt=0"`
		RoleID      int16     `json:"role_id" validate:"required gt=0"`
		Nickname    string    `json:"nickname" validate:"required"`
		Name        string    `json:"name" validate:"required"`
		LastName    string    `json:"last_name" validate:"required"`
		DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddPlayer(c.Request().Context(), req.TeamID, req.RoleID,
			req.Nickname, req.Name, req.LastName, req.DateOfBirth); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeletePlayer(svc services.PlayersServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		if err := svc.DeletePlayer(c.Request().Context(), playerID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetPlayersByMatchKDA(svc services.PlayersServiceInterface) echo.HandlerFunc {
	type request struct {
		KDA float64 `json:"kda" validate:"required gte=0"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		players, err := svc.FindPlayersByMatchKDA(c.Request().Context(), req.KDA)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, players)
	}
}
func UpdatePlayerNickname(svc services.PlayersServiceInterface) echo.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname" validate:"required"`
	}
	return func(c echo.Context) error {
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.UpdatePlayerNickname(c.Request().Context(), playerID, req.Nickname); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}

// Stats
func GetStatsByPlayer(svc services.StatsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		stats, err := svc.GetByPlayer(c.Request().Context(), playerID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, stats)
	}
}
func AddStatistic(svc services.StatsServiceInterface) echo.HandlerFunc {
	type request struct {
		Kills         int16   `json:"kills" validate:"required gte=0"`
		Deaths        int16   `json:"deaths" validate:"required gte=0"`
		Assists       int16   `json:"assists" validate:"required gte=0"`
		Creeps        int16   `json:"creeps" validate:"required gte=0"`
		Denieds       int16   `json:"denieds" validate:"required gte=0"`
		GoldPerMinute float64 `json:"gold_per_minute" validate:"required gt=0"`
		NetWorth      float64 `json:"net_worth" validate:"required gte=0"`
	}
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddStatistic(c.Request().Context(), matchID, playerID, req.Kills,
			req.Deaths, req.Assists, req.Creeps, req.Denieds, req.GoldPerMinute, req.NetWorth); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteStatsByMatch(svc services.StatsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		if err := svc.DeleteByMatch(c.Request().Context(), matchID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func DeleteStatsByID(svc services.StatsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		playerID, err := strconv.Atoi(c.Param("playerID"))
		if err != nil || playerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid playerID"})
		}
		if err := svc.DeleteByID(c.Request().Context(), matchID, playerID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetStatsByMatch(svc services.StatsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchID, err := strconv.Atoi(c.Param("matchID"))
		if err != nil || matchID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid matchID"})
		}
		stats, err := svc.GetByMatch(c.Request().Context(), matchID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, stats)
	}
}

// Stage
func GetStagesByTournament(svc services.StageServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		stages, err := svc.GetByTournament(c.Request().Context(), tournamentID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, stages)
	}
}
func GetStageByID(svc services.StageServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		stageID, err := strconv.Atoi(c.Param("stageID"))
		if err != nil || stageID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid stageID"})
		}
		stage, err := svc.GetByID(c.Request().Context(), stageID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, stage)
	}
}
func AddStage(svc services.StageServiceInterface) echo.HandlerFunc {
	type request struct {
		StageName models.StageType `json:"stage_name" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		if err := svc.AddStage(c.Request().Context(), tournamentID, req.StageName); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteStage(svc services.StageServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		stageID, err := strconv.Atoi(c.Param("stageID"))
		if err != nil || stageID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid stageID"})
		}
		if err := svc.Delete(c.Request().Context(), stageID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}

// Teams
func GetTeamsByTournament(svc services.TeamsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tournamentID, err := strconv.ParseInt(c.Param("tournamentID"), 10, 64)
		if err != nil || tournamentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid tournamentID"})
		}
		teams, err := svc.FindTeamByTournament(c.Request().Context(), tournamentID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, teams)
	}
}
func GetAllTeams(svc services.TeamsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teams, err := svc.GetAllTeams(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, teams)
	}
}
func GetTeamsWinrates(svc services.TeamsServiceInterface) echo.HandlerFunc {
	type request struct {
		Winrate float64 `json:"winrate" validate:"required gt=0 lt=100"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		teams, err := svc.GetTeamsWinrates(c.Request().Context(), req.Winrate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, teams)
	}
}
func GetTeamsFaceToFace(svc services.TeamsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teams, err := svc.GetTeamsFaceToFace(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, teams)
	}
}
func AddTeam(svc services.TeamsServiceInterface) echo.HandlerFunc {
	type request struct {
		Country        string    `json:"country" validate:"required "`
		Title          string    `json:"title" validate:"required"`
		FoundationDate time.Time `json:"foundation_date" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddTeam(c.Request().Context(), req.Country, req.Title, req.FoundationDate); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteTeam(svc services.TeamsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID, err := strconv.Atoi(c.Param("teamID"))
		if err != nil || teamID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid teamID"})
		}
		if err := svc.DeleteTeam(c.Request().Context(), teamID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
