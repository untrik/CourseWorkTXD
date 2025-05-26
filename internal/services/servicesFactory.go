package services

import "github.com/untrik/CourseWorkTXD/internal/repositories"

type ServicesFactory struct {
	Matches     MatchesServiceInterface
	Players     PlayersServiceInterface
	Teams       TeamsServiceInterface
	Tournaments TournamentsServiceInterface
	Stage       StageServiceInterface
	Participant ParticipantsServiceInterface
	Result      MatchesResultServiceInterface
	Stats       StatsServiceInterface
}

func NewServicesFactory(repoFactory *repositories.RepoFactory) *ServicesFactory {
	return &ServicesFactory{
		Matches:     NewMatchesService(repoFactory.Matches, repoFactory.Teams, repoFactory.Tournaments, repoFactory.Stage),
		Players:     NewPlayersService(repoFactory.Players, repoFactory.Teams),
		Teams:       NewTeamsService(repoFactory.Teams, repoFactory.Tournaments),
		Tournaments: NewTournamentsService(repoFactory.Tournaments),
		Stage:       NewStageService(repoFactory.Stage, repoFactory.Tournaments),
		Participant: NewParticipantsService(repoFactory.Participants, repoFactory.Teams, repoFactory.Matches, repoFactory.Tournaments),
		Result:      NewMatchesResultService(repoFactory.Result, repoFactory.Matches, repoFactory.Teams),
		Stats:       NewStatsService(repoFactory.Stats, repoFactory.Matches, repoFactory.Players)}

}
