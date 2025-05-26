package repositories

import (
	"gorm.io/gorm"
)

type RepoFactory struct {
	Matches      MatchesRepoInterface
	Players      PlayersRepoInterface
	Teams        TeamRepoInterface
	Tournaments  TournamentsRepoInterface
	Stage        StageRepoInterface
	Participants ParticipantsRepoInterface
	Result       MatchesResultRepoInterface
	Stats        StatsRepoInterface
}

func NewRepoFactory(db *gorm.DB) *RepoFactory {
	return &RepoFactory{
		Matches:      NewMatchesRepo(db),
		Players:      NewPlayersRepo(db),
		Teams:        NewTeamsRepo(db),
		Tournaments:  NewTournamentsRepo(db),
		Stage:        NewStageRepo(db),
		Participants: NewParticipantsRepo(db),
		Result:       NewMatchesResultRepo(db),
		Stats:        NewStatsRepo(db)}
}
