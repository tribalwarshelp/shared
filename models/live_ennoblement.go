package models

import "time"

type LiveEnnoblement struct {
	VillageID  int       `gqlgen:"-" json:"villageID"`
	Village    *Village  `gqlgen:"-" json:"village,omitempty"`
	NewOwnerID int       `gqlgen:"-" json:"newOwnerID"`
	NewOwner   *Player   `gqlgen:"-" json:"newOwner,omitempty"`
	OldOwnerID int       `gqlgen:"-" json:"oldOwnerID"`
	OldOwner   *Player   `gqlgen:"-" json:"oldOwner,omitempty"`
	EnnobledAt time.Time `pg:"default:now(),use_zero" json:"ennobledAt" gqlgen:"ennobledAt"`
}
