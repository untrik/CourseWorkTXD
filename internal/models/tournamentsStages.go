package models

type StageType string

const (
	GroupStage         StageType = "group_stage"
	UpBrQuarterfinals  StageType = "upper_bracket_quarterfinals"
	UpBrSemifinals     StageType = "upper_bracket_semifinals"
	UpBrFinal          StageType = "upper_bracket_final"
	GrandFinal         StageType = "grand_final"
	LowBrRound         StageType = "lower_bracket_round_1"
	LowBrQuarterfinals StageType = "lower_bracket_quarterfinals"
	LowBrSemifinals    StageType = "lower_bracket_semifinal"
	LowBrFinal         StageType = "lower_bracket_final"
	Replays            StageType = "replays"
)

func (st StageType) IsValid() bool {
	switch st {
	case GroupStage,
		UpBrQuarterfinals,
		UpBrSemifinals,
		UpBrFinal,
		GrandFinal,
		LowBrRound,
		LowBrQuarterfinals,
		LowBrSemifinals,
		LowBrFinal,
		Replays:
		return true
	default:
		return false
	}
}

type TournamentsStages struct {
	StageID      int         `json:"stage_id" gorm:"primaryKey;autoIncrement;not null"`
	TournamentID int64       `json:"tournament_id" gorm:"not null"`
	Tournament   Tournaments `gorm:"foreignKey:TournamentID"`
	StageName    StageType   `json:"stage_name" gorm:"not null;size:200"`
}
