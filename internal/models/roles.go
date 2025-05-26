package models

type RolesTitle string

const (
	Carry       RolesTitle = "carry"
	Mid         RolesTitle = "mid"
	Offlane     RolesTitle = "offlane"
	SoftSupport RolesTitle = "soft_support"
	HardSupport RolesTitle = "hard_support"
)

func (rt RolesTitle) IsValid() bool {
	switch rt {
	case Carry,
		Mid,
		Offlane,
		SoftSupport,
		HardSupport:
		return true
	default:
		return false
	}
}

type Roles struct {
	ID    int16      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title RolesTitle `json:"title" gorm:"size:200;not null;unique"`
}
