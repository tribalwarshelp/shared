package models

import "time"

type Ennoblement struct {
	VillageID  int       `gqlgen:"-" json:"villageID"`
	Village    *Village  `gqlgen:"-" json:"village"`
	NewOwnerID int       `gqlgen:"-" json:"newOwnerID"`
	NewOwner   *Player   `gqlgen:"-" json:"newOwner"`
	OldOwnerID int       `gqlgen:"-" json:"oldOwnerID"`
	OldOwner   *Player   `gqlgen:"-" json:"oldOwner"`
	EnnobledAt time.Time `json:"ennobledAt" gqlgen:"ennobledAt"`
}
